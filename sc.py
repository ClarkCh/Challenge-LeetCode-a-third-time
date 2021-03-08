import re


def init():
    with open(fileName, "r+") as file:
        lines = file.readlines()
#     lines = """
# /**
#  * Definition for singly-linked list.
#  * type ListNode struct {
#  *     Val int
#  *     Next *ListNode
#  * }
#  */
#  """
    start = 0
    res = []
    for line in lines:
        res.append(line)
        if start:
            res.pop()
            res.append(line[3:])
            if " */" == line[:3]:
                res.pop()
                start = 0
        if " * Definition for" in line:
            res.pop()
            res.pop()
            start = 1
    with open(fileName, "w") as file:
        file.write("".join(res))


    # res = re.findall("/\*\*\n \* Definition.*?\n( \* .*?)*?", lines, re.DOTALL)
    # print(res)

if __name__ == "__main__":
    fileName = "2两数相加.go"
    init()


