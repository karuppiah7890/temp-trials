# Story

Interesting extra ways to find out decimal to binary. I knew only one previously

https://duckduckgo.com/?t=ffab&q=decimal+to+binary+conversion+process&ia=web

https://www.tutorialspoint.com/how-to-convert-decimal-to-binary

---

I'm planning to simply divide the number by 2 till I get the binary digits and also try to find the bigger binary gap :D

So, for example, for 1041, it would be

1041 / 2 = 520 + 1/2

So, 1 is the remainder. I start keeping a tab of the 0s that come next as there's a 1 and I want to count all the 0s till the next 1 if there's any. So, when I encounter the next 1, I store the count of 0s in between the two 1s as a binary gap and keep a tab of the biggest binary gap which is the output

520 / 2 = 260 + 0/2

So, 0 is the remainder

I can go on like this till I end up with nothing to divide

The actual representation of 1041 in binary is 10000010001 , in the above method, I get the output backwards, like this - 10001000001 and it has to be reversed to get the actual binary representation

But the reverse or actual doesn't matter when I'm only trying to find the binary gap between two 1s in the representation - which doesn't change even if the same representation is reversed. So, I don't have to do any reverse :D

Also, I don't have to wait till the end to get a representation and then traverse it to get the largest binary gap, instead I can get it on the go while dividing the value and getting the binary digits

I think the time complexity here is O(log N) as we keep dividing the value by 2 and it goes on till we have nothing to divide, so for 1041, it will be around 10 as 2^10 is 1024

---

https://duckduckgo.com/?t=ffab&q=golang+division&ia=web&iax=qa

Golang actually has a nice `bits` package, hmm

https://golang.org/pkg/math/

https://golang.org/pkg/math/bits/

But I don't get what `hi` and `lo` mean, hmm. Wondering if I should use this for an optimal solution ;) Let's see

Also, for division in Golang, looks like I can't get quotient and remainder in one go. Instead I have to do the operation twice - once to get quotient, another to get remainder. Hmm. So, it's like O(log N + log N) = O(2log N)?


