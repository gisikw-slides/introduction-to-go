void primes(int m, int t, int c) {
  ((t / m) <= 1) ? primes(m,t+1,c) : !(t % m) ? primes(m,t+1, t % m) : 
  ((t % m)==(t / m) && !c) ? (printf("%d\t",(t / m)), primes(m,t+1,c)) : 
  ((t % m)> 1 && (t % m) < (t / m)) ? primes(m,t+1,c + !((t / m) % (t % m))) : 
  (t < m * m) ? primes(m,t+1,c) : 0;
}
