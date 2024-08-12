class Queue:
    def __init__(self):
        self.queue=[]
    #add queue
    def enqueue(self, item):
        self.queue.append(item)
    #remove item from queue
    def dequeue(self):
        if len(self.queue)<1:
            return None
        return self.queue.pop(0)
    #display queue
    def display(self):
        print(self.queue)
    #size
    def size(self):
        print(len(self.queue))

q=Queue()
q.enqueue(1)
q.enqueue(3)
q.enqueue(7)
q.display()
q.dequeue()
q.display()
q.size()