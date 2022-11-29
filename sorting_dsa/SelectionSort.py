def selectionsort(A):
    n = len(A)
    for i in range(n-1):
        position = i
        for j in range(i+1, n):
            if A[j] < A[position]:
                position = j
        temp = A[i]
        A[i] = A[position]
        A[position] = temp


L = [3, 5, 8, 9, 6, 2]
print('Before sorting, Original list:',L)
selectionsort(A)
print('After sorting, Sorted list:',L)
