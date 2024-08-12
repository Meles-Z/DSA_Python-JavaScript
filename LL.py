class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None
    
    def insert_at_beginning(self, data):
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node
    
    def insert_at_middle_after(self, prevnode, data):
        curr=self.head
        while curr and curr.data !=prevnode:
            curr=curr.next

        if curr is None:
            print("To insert at the middle prevnode must exits")
            return
        
        new_node=Node(data)
        new_node.next=curr.next
        curr.next=new_node
     
    def insert_at_middle_before(self, nextNode, data):
        curr=self.head
        while curr and curr.next and curr.next.data !=nextNode:
            curr=curr.next
        if curr is None and curr.next is None:
            print("Node with value", nextNode, "not found")
            return
        new_node=Node(data)
        new_node.next=curr.next
        curr.next=new_node

        
        
    def insert_at_end(self, data):
        new_node=Node(data)
        if self.head is None:
            self.head=new_node
            return
        curr=self.head
        while curr.next:
            curr=curr.next
        curr.next=new_node
    
    def get_numbers(self):
        numbers = []
        curr = self.head
        while curr:
            numbers.append(curr.data)
            curr = curr.next
        return numbers

ll = LinkedList()
ll.insert_at_beginning(2)
ll.insert_at_beginning(3)
ll.insert_at_beginning(6)
ll.insert_at_end(9)
ll.insert_at_middle_after(3, 12)
ll.insert_at_middle_before(3,20)
numbers = ll.get_numbers()
print(numbers)  
