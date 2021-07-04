An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.

Write a function:

    func Solution(A []int, K int) []int

that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given
    A = [3, 8, 9, 7, 6]
    K = 3

the function should return [9, 7, 6, 3, 8]. Three rotations were made:
    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]

For another example, given
    A = [0, 0, 0]
    K = 1

the function should return [0, 0, 0]

Given
    A = [1, 2, 3, 4]
    K = 4

the function should return [1, 2, 3, 4]

Assume that:

- N and K are integers within the range [0..100];
- Each element of array A is an integer within the range [−1,000..1,000].

In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.

---

Possible optimizations -
- When K = length(A), then it's a full rotation leading to the result as A, so we don't even have to do a rotation here
- When the value of K is too high relative to the length(A). For example, if length(A) = 10 and K = 8, that's 8 right rotations. Instead, one could do 2 left rotations. So, we can can define a K' which can be defined as a left rotation equivalent to the K right rotations. K' = length(A) - K. And then one can do minimum(K, K') to see which is the smallest and then accordingly choose to do K right rotations or an equivalent K' left rotations, whichever is least effort / processing :) Also, if you notice, in the special case of K = length(A), K' = length(A) - K = 0 and minimum(K, K') = K' = 0, so we don't do any rotations basically ;) :D, so it kind of does what we mentioned in the first optimization

---

https://duckduckgo.com/?t=ffab&q=golang+array+operations&ia=web

https://tour.golang.org/moretypes/7

https://tour.golang.org/moretypes/9

https://tour.golang.org/moretypes/10

https://tour.golang.org/moretypes/11

https://tour.golang.org/moretypes/12

https://tour.golang.org/moretypes/13

https://tour.golang.org/moretypes/14

https://tour.golang.org/moretypes/15

---

Oops, I didn't consider the case where K can actually be greater than length(A). I always assumed it will be lesser, though it can be greater. Hmm

The following test case failed - 

Array: []
K: 1
Expected Output: []

---

Test results - PDF

[cyclic-rotation-test-results-codility.pdf](./cyclic-rotation-test-results-codility.pdf)

---

TODO:
- Fix runtime errors - slice bounds out of range errors
- Fix assumption that K <= length(A)
