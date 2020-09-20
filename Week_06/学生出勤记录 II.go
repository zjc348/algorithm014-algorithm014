func checkRecord(n int) int {
    M := 1000000007;
    a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 := 1, 0, 0, 0, 0, 0;
    for i := 0; i <= n; i++ {
        a0l0_ := (a0l0 + a0l1 + a0l2) % M;
        a0l2 = a0l1;
        a0l1 = a0l0;
        a0l0 = a0l0_;
        a1l0_ := (a0l0 + a1l0 + a1l1 + a1l2) % M;
        a1l2 = a1l1;
        a1l1 = a1l0;
        a1l0 = a1l0_;
    }
    return a1l0;
}