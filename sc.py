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
    rec = []
    for line in lines:
        rec.append(line)
        if start:
            rec.pop()
            rec.append(line[3:])
            if " */" == line[:3]:
                rec.pop()
                start = 0
        if " * Definition for" in line:
            rec.pop()
            rec.pop()
            start = 1
    with open(fileName, "w") as file:
        file.write("".join(rec))


    # rec = re.findall("/\*\*\n \* Definition.*?\n( \* .*?)*?", lines, re.DOTALL)
    # print(rec)

if __name__ == "__main__":
    fileName = "2两数相加.go"
    init()


