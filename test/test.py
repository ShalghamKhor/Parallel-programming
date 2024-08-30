from queue import Queue
from threading import Semaphore, Thread

class IQS:

    def __init__(self):
        self.Q = Queue()
        self.qsem = Semaphore(0)

    
    def _sort(self):
        while True:
            self.qsem.acquire()
            lo, hi = self.Q.get()
            if lo < hi:
                mid = self._partition(lo, hi) # Unchanged11
                self.Q.put((lo, mid - 1))
                self.Q.put((mid + 1, hi))
                self.qsem.release(2)

    def sort(self, l):
        self.lst = l
        self.Q.put((0, len(self.lst) - 1))
        self.qsem.release()
        ts = [Thread(target=self._sort) for _ in range(self.nt)]

        for t in ts:
            t.start()
        
        for t in ts:
            t.join()