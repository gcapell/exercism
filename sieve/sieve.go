package sieve

func Sieve(n int) []int {
	switch n {
		case 0,1,2: return []int{2}
		case 3: return []int{2,3}
	}
	
	primes := make([]int, 0, 170)
	primes = append(primes, 2, 3)
	sieve := make([]bool, n)
	p := 5 
	for p <= n {
		if !sieve[p] {
			primes = append(primes, p)
			filter(sieve, p)
		}
		p += 2
		if !sieve[p] {
			primes = append(primes, p)
			filter(sieve, p)
		}
		p += 4
	}
	// We might have added an extra 
	if primes[len(primes)-1] > n {
		primes = primes[:len(primes)-1]
	}
	return primes
}

func filter(sieve []bool, n int) {
	for j := n; j< len(sieve); j += n{
		sieve[j]=true
	}
}