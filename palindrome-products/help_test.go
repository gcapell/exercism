package palindrome

func TestPalLess(t *testing.T) {
tests:
	for _, test := range []struct {
		start  int
		expect []int
	}{
		{999, []int{999, 989, 979, 969, 959, 949, 939, 929, 919, 909, 898, 888}},
		{998, []int{989, 979}},
		{10, []int{9, 8, 7}},
		{1000, []int{999, 989, 979, 969, 959}},
		{10000, []int{9999, 9889, 9779}},
	} {
		gen := palLess(test.start)
		for pos, want := range test.expect {
			got := gen()
			if got != want {
				t.Errorf("palLess(%d)[%d] = %v, want %v", test.start, pos, got, want)
				continue tests
			}
		}
	}
}

func TestPalGreater(t *testing.T) {
tests:
	for _, test := range []struct {
		start  int
		expect []int
	}{
		{999, []int{999, 1001, 1111, 1221}},
		{1, []int{1,2,3,4,5,6,7,8,9, 11, 22, 33,}},
		{80, []int{88, 99, 101, 111, 121, 131,}},
		{190, []int{191, 202, }},
		{999, []int{999, 1001, 1111,1221,}},
		{9998, []int{9999, 10001, 10101, 10201}},
	} {
		gen := palGreater(test.start)
		for pos, want := range test.expect {
			got := gen()
			if got != want {
				t.Fatalf("palGreater(%d)[%d] = %v, want %v", test.start, pos, got, want)
				continue tests
			}
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	for _, test := range []struct {
		n    int
		want []Factor
	}{
		{999, []Factor{{3, 3}, {37, 1}}},
		{4, []Factor{{2, 2}}},
	} {
		fs, _ := primeFactors(test.n)
		if !reflect.DeepEqual(fs, test.want) {
			t.Errorf("primeFactors(%d)=%v, want %v", test.n, fs, test.want)
		}
	}
}

func TestGroup(t *testing.T) {
	for _, test := range []struct {
		factors    []Factor
		small, big	int
		fmin, fmax int
		want [][2]int
	}{
		{
			factors: []Factor { {2,2}, {3,1}, {5,2}, {7,1}},
			small:1,
			big:1,
			fmin:10,
			fmax:100,
			want: [][2]int {{21, 100}, {25, 84}, {28, 75}, {30, 70}, {35 ,60}, {42, 50},},
		},
	} {
		got := group(test.factors, test.small, test.big, test.fmin, test.fmax)
		sort.Sort(ByFirst(got))
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("group(%v,%d,%d,%d,%d)=%v, want %v", 
				test.factors, test.small, test.big, test.fmin, test.fmax, 
				got, test.want)
		}
	}
}
