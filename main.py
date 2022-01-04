# -*- coding: utf-8 -*-

import json
import os
import time

import requests

# ~~~~~~~~~~~~以下是无需修改的参数~~~~~~~~~~~~~~~~·
requests.packages.urllib3.disable_warnings()  # 为了避免弹出一万个warning，which is caused by 非验证的get请求

leetcode_url = 'https://leetcode-cn.com/'

sign_in_url = leetcode_url + 'accounts/login/'
submissions_url = leetcode_url + 'submissions/'

config_path = "./configMy.json"
user_agent = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) ' \
             'Chrome/33.0.1750.152 Safari/537.36 '

with open(config_path, "r") as f:  # 读取用户名，密码，本地存储目录
    config = json.loads(f.read())
    USERNAME = config['username']
    PASSWORD = config['password']
    OUTPUT_DIR = config['outputDir']
    TIME_CONTROL = 3600 * config['time']

FILE_FORMAT = {"C++": ".cpp", "Python3": ".py", "Python": ".py", "MySQL": ".sql", "Go": ".go", "Java": ".java",
               "C": ".c", "JavaScript": ".js", "PHP": ".php", "C#": ".cs", "Ruby": ".rb", "Swift": ".swift",
               "Scala": ".scl", "Kotlin": ".kt", "Rust": ".rs"}
# ~~~~~~~~~~~~以上是无需修改的参数~~~~~~~~~~~~~~~~·

# ~~~~~~~~~~~~以下是可以修改的参数~~~~~~~~~~~~~~~~·
START_PAGE = 0  # 从哪一页submission开始爬起，0是最新的一页
SLEEP_TIME = 5  # in second，登录失败时的休眠时间


# ~~~~~~~~~~~~以上是可以修改的参数~~~~~~~~~~~~~~~~·

# 登陆
def login(username, password):  # 本函数修改自https://gist.github.com/fyears/487fc702ba814f0da367a17a2379e8ba，感谢@fyears
    client = requests.session()
    client.encoding = "utf-8"

    while True:
        try:
            client.get(sign_in_url, verify=False)
            login_data = {
                'login': username,
                'password': password
            }
            result = client.post(sign_in_url, data=login_data, headers=dict(Referer=sign_in_url))
            if result.ok:
                print("Login successfully!")
                break
        except:
            print("登录失败！")
            time.sleep(SLEEP_TIME)
    return client


# 获取最新题解TimeStamp与url id的对应
def getTimeStamp(client):
    detailed_question_url = "https://leetcode-cn.com/graphql"
    question_headers = {'User-Agent': user_agent,
                        'Connection': 'keep-alive',
                        'Content-Type': 'application/json',
                        'Referer': 'https://leetcode-cn.com/problems/'}
    count = 0
    cnt = 0
    while True:
        cur_time = time.time()
        params = {'operationName': "userProfileQuestions",
                  'variables': {"status": "ACCEPTED",
                                "sortOrder": "DESCENDING",
                                "sortField": "LAST_SUBMITTED_AT",
                                "first": 20,
                                "difficulty": [],
                                "skip": count
                                },
                  'query':
                      '''
            query userProfileQuestions($status: StatusFilterEnum!, $skip: Int!, $first: Int!, $sortField: SortFieldEnum!, $sortOrder: SortingOrderEnum!, $keyword: String, $difficulty: [DifficultyEnum!]) {
            userProfileQuestions(status: $status, skip: $skip, first: $first, sortField: $sortField, sortOrder: $sortOrder, keyword: $keyword, difficulty: $difficulty) {
                totalNum
                questions {
                translatedTitle
                frontendId
                titleSlug
                title
                difficulty
                lastSubmittedAt
                numSubmitted
                lastSubmissionSrc {
                    sourceType
                    ... on SubmissionSrcLeetbookNode {
                    slug
                    title
                    pageId
                    __typename
                    }
                __typename
                }
                __typename
                }
            __typename
            }
            } 
            '''
                  }
        json_data = json.dumps(params).encode('utf8')
        response = client.post(detailed_question_url, data=json_data, headers=question_headers, timeout=10,
                               verify=False)
        if response.status_code != 200:
            print("总共", cnt, "道题目")
            time.sleep(1)
            return
        content = response.json()
        if len(content['data']['userProfileQuestions']['questions']) == 0:
            print("总共", cnt, "道题目")
            time.sleep(1)
            return
        for question in content['data']['userProfileQuestions']['questions']:
            if cur_time - question['lastSubmittedAt'] > TIME_CONTROL:  # 时间管理，本行代表只记录最近的TIME_CONTROL天内的提交记录
                print("总共", cnt, "道题目")
                time.sleep(1)
                return
            frontendId[question['lastSubmittedAt']] = question['frontendId']
            titleSlug[question['lastSubmittedAt']] = question['titleSlug']
            print(question['frontendId'], '\t', question['translatedTitle'])
            cnt += 1
        count += 20
        time.sleep(1)


def scraping(client):
    page_num = START_PAGE
    visited = set()
    cnt = 1
    while True:
        print("Now for page:", str(page_num))
        submissions_url = "https://leetcode-cn.com/api/submissions/?offset=" + str(page_num) + "&limit=40&lastkey="

        html = client.get(submissions_url, verify=False)
        html = json.loads(html.text)
        cur_time = time.time()
        if not html.get("submissions_dump"):
            print("下载完成")
            break

        for submission in html["submissions_dump"]:
            submission_status = submission['status_display']
            problem_title = submission['title'].replace(" ", "")
            submission_language = submission['lang']

            if submission_status != "Accepted":
                continue

            if cur_time - submission['timestamp'] > TIME_CONTROL:  # 时间管理，本行代表只记录最近的TIME_CONTROL天内的提交记录
                print("完成所需时间的下载")
                return
            try:
                if submission['timestamp'] in frontendId:
                    # 下载题目
                    url = titleSlug[submission['timestamp']]
                    tags = get_question_detail(url)
                    # 下载题解
                    problem_id = frontendId[submission['timestamp']]
                    if problem_id + submission_language not in visited:
                        visited.add(problem_id + submission_language)  # 保障每道题只记录每种语言最新的AC解
                        full_paths = generatePath(problem_id, problem_title, submission_language, tags)
                        print(cnt, ": ", problem_id, problem_title, submission_language)
                        cnt+=1
                        del frontendId[submission['timestamp']]
                        time.sleep(1)
                        code = downloadCode(submission, client)
                        for full_path in full_paths:
                            with open(full_path, "w") as f:  # 开始写到本地
                                f.write(code)
            except FileNotFoundError as e:
                print("文件夹不存在")

        time.sleep(1)
        page_num += 40

# 下载题目
def get_question_detail(simple_url):
    session = requests.Session()
    detailed_question_url = "https://leetcode-cn.com/graphql"
    question_headers = {'User-Agent': user_agent,
                        'Connection': 'keep-alive',
                        'Content-Type': 'application/json',
                        'Referer': 'https://leetcode-cn.com/problems/' + simple_url}
    params = {'operationName': "getQuestionDetail",
              'variables': {'titleSlug': simple_url},
              'query':
                  '''
                  query getQuestionDetail($titleSlug: String!) {
                            question(titleSlug: $titleSlug) {
                                questionId
                                questionFrontendId
                                titleSlug
                                title
                                translatedTitle 
                                difficulty
                                categoryTitle
                                similarQuestions
                                topicTags {
                                    name
                                    slug
                                    translatedName
                                }
                                content
                                translatedContent
                            }
                  }
                  '''
              }
    json_data = json.dumps(params).encode('utf8')
    response = session.post(detailed_question_url, data=json_data, headers=question_headers, timeout=10)
    content = response.json()
    if content['data']['question']['translatedTitle'] is None:
        session.close()
        return
    process_writing_question(content)
    session.close()
    return content['data']['question']['topicTags']


# 保存题目
def process_writing_question(content):
    question_url = "https://leetcode-cn.com/problems/"
    tag_url = "https://leetcode-cn.com/tag/"

    frontend_id = content['data']['question']['questionFrontendId']
    title_cn = content['data']['question']['translatedTitle']
    content_cn = content['data']['question']['translatedContent']
    simple_url = content['data']['question']['titleSlug']
    tags = content['data']['question']['topicTags']

    full_paths = generatePath(frontend_id, title_cn.replace(" ", ""), "md", tags)
    for full_path in full_paths:
        f_cn = open(full_path, 'w', encoding='UTF-8')
        f_cn.write("# [{}. {}]({})\n".format(frontend_id, title_cn, question_url + simple_url))

        f_cn.write(content_cn if content_cn else "")

        if len(tags) > 0:
            f_cn.write("\n**标签:**  ")
            for tag in tags:
                f_cn.write("[{}]({}) ".format(tag['translatedName'], tag_url + tag['slug']))

        f_cn.close()

def downloadCode(submission, client):
    # print(submission)
    headers = {
        'Connection': 'keep-alive',
        'Content-Type': 'application/json',
        'User-Agent': user_agent
    }
    param = {'operationName': "mySubmissionDetail", "variables": {"id": submission["id"]},
             'query': "query mySubmissionDetail($id: ID\u0021) {  submissionDetail(submissionId: $id) {    id    code    runtime    memory    statusDisplay    timestamp    lang    passedTestCaseCnt    totalTestCaseCnt    sourceUrl    question {      titleSlug      title      translatedTitle      questionId      __typename    }    ... on GeneralSubmissionNode {      outputDetail {        codeOutput        expectedOutput        input        compileError        runtimeError        lastTestcase        __typename      }      __typename    }    __typename  }}"
             }

    param_json = json.dumps(param).encode("utf-8")
    response = client.post("https://leetcode-cn.com/graphql/", data=param_json, headers=headers)
    submission_details = response.json()["data"]["submissionDetail"]

    return submission_details["code"]


def generatePath(problem_id, problem_title, submission_language, tags):
    if not os.path.exists(OUTPUT_DIR):
        os.mkdir(OUTPUT_DIR)
    if not os.path.exists(OUTPUT_DIR + "/全部"):
        os.mkdir(OUTPUT_DIR + "/全部")
    if not os.path.exists(OUTPUT_DIR + "/标签"):
        os.mkdir(OUTPUT_DIR + "/标签")

    if problem_id[0].isdigit():  # 如果题目是传统的数字题号
        newPath = OUTPUT_DIR + "/全部/" + '{:0=4}'.format(int(problem_id)) + "." + problem_title  # 存放的文件夹名
        if submission_language == "md":
            filename = "README.md"  # 存放的文件名
        else:
            filename = '{:0=4}'.format(int(problem_id)) + "-" + problem_title + FILE_FORMAT[submission_language]  # 存放的文件名
    else:  # 如果题目是新的面试题
        newPath = OUTPUT_DIR + "/全部/" + problem_id + "." + problem_title
        if submission_language == "md":
            filename = "README.md"  # 存放的文件名
        else:
            filename = problem_id + "-" + problem_title + FILE_FORMAT[submission_language]  # 存放的文件名

    if not os.path.exists(newPath):
        os.mkdir(newPath)
    full_path = []
    full_path.append(os.path.join(newPath, filename))  # 把文件夹和文件组合成新的地址

    if len(tags) > 0:
        for tag in tags:
            tag = tag['translatedName']
            if not os.path.exists(OUTPUT_DIR + "/标签/" + tag):
                os.mkdir(OUTPUT_DIR + "/标签/" + tag)
            if problem_id[0].isdigit():  # 如果题目是传统的数字题号
                newPath = OUTPUT_DIR + "/标签/" + tag + "/" + '{:0=4}'.format(int(problem_id)) + "." + problem_title  # 存放的文件夹名
                if submission_language == "md":
                    filename = "README.md"  # 存放的文件名
                else:
                    filename = '{:0=4}'.format(int(problem_id)) + "-" + problem_title + FILE_FORMAT[submission_language]  # 存放的文件名
            else:  # 如果题目是新的面试题
                newPath = OUTPUT_DIR + "/标签/" + tag + "/" + problem_id + "." + problem_title
                if submission_language == "md":
                    filename = "README.md"  # 存放的文件名
                else:
                    filename = problem_id + "-" + problem_title + FILE_FORMAT[submission_language]  # 存放的文件名
            if not os.path.exists(newPath):
                os.mkdir(newPath)
            full_path.append(os.path.join(newPath, filename))  # 把文件夹和文件组合成新的地址
    return full_path


def gitPush():
    today = time.strftime('%Y-%m-%d', time.localtime(time.time()))
    os.chdir(OUTPUT_DIR)
    instructions = ["git add .", "git status", "git commit -m \"" + today + "\"", "git push -u origin master"]
    for ins in instructions:
        os.system(ins)
        print("~~~~~~~~~~~~~" + ins + " finished! ~~~~~~~~")


def main():
    print('Login')
    client = login(USERNAME, PASSWORD)
    print('获取 时间戳')
    getTimeStamp(client)
    print('开始下载')
    scraping(client)
    gitPush()


if __name__ == '__main__':
    frontendId = {}
    titleSlug = {}
    main()
    print(frontendId)
