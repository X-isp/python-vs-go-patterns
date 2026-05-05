import time
import random


def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]


data = [random.randint(0, 10000) for _ in range(2000)]
data_copy = data.copy()

start = time.time()
bubble_sort(data_copy)
end = time.time()

print(f"Python: {end - start:.6f}")