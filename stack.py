
#lets test stack

def create_stack():
    stack=[]
    return stack
def is_impty(stack):
    if len(stack)==0:
        print("stack is empty")
def push(stack, item):
    stack.append(item)
def pop(stack):
    if is_impty(stack):
        return
    return stack.pop()

s=create_stack()
push(s, 1)
push(s, 2)
push(s,3)
print(s)
pop(s)
print(s)



class Solution(object):
    def is_palindrome(self, x, target):
        for i in range(len(x)):
            for j in range(i+1, len(x)):
                if x[i]+x[j]==target:
                    return[i,j]
        return[]