# race_condition_in_go

## Problem 

Here we tried increasing the value of variable x by 1000 parallel go routine
initiating from 0 so ideal value should be 1000 

but....

due to race condition amongst the go routines the value is not reaching to 1000

## Solution

We can solve the race condition issue via 2 approaches 

### 1. Using Mutex :

Go provide Mutex out of the box which ensures that the critical section [ Here it is increasing variable value] can be accessed only by 1 of the go-routine 
this is implemented in fix/using_mutex.go

### 2. Using Channels :

We have used bufferred bool channel to wait on the process till the channel is not emptied 
and once the channel is emptied then only the next process can start the process , this helped
in ensuring critical section can only be accessed once 


Referece is taken from : (https://golangbot.com/mutex/)