package solve_test

/*
func TestSpecifics(t *testing.T) {
	solve.SetTestTimeout()
	// go decided that it should run tests in this directory.
	os.Chdir(`..`)

	testCases := []struct {
		id   string
		iter model.Iterator
	}{{
		iter: 8,
		id:   `7,437,536`,
	}}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.iter.String()+` `+tc.id, func(t *testing.T) {
			sr, err := fetch.ReadID(tc.iter, tc.id)
			if err != nil || sr.Input == nil {
				t.Logf("Error fetching input: %q", err)
				t.Fail()
			}

			ns := sr.Input.ToNodes()
			sol, err := solve.FromNodesWithTimeout(
				tc.iter.GetWidth(),
				tc.iter.GetHeight(),
				ns,
				50*time.Second,
			)
			if err != nil {
				t.Logf("Error fetching input: %q", err)
				t.Fail()
			}
			if sol.ToAnswer() != sr.Answer {
				t.Logf("Incorrect Answer\n")
				t.Logf("Exp: %q\n", sr.Answer)
				t.Logf("Act: %q\n", sol.ToAnswer())
				t.Logf("Board:\n%s\n\n", sol.Pretty(ns))
				t.Fail()
			}
		})
	}
}

func TestAccuracy(t *testing.T) {
	fetch.DisableHTTPCalls()
	solve.SetTestTimeout()
	// go decided that it should run tests in this directory.
	os.Chdir(`..`)
	max := model.MaxIterator

	for iter := model.MinIterator; iter <= max; iter++ {
		t.Run(iter.String(), func(t *testing.T) {
			srs, err := fetch.ReadN(iter, 100)
			if err != nil {
				t.Logf("Error fetching input: %q", err)
				t.Fail()
			}
			if len(srs) == 0 {
				t.Skip()
			}

			for _, sr := range srs {
				sr := sr
				if sr.Answer == `` {
					t.Logf("Unknown answer: %+v", sr)
					t.Fail()
				}
				t.Run(sr.Input.ID, func(t *testing.T) {
					ns := sr.Input.ToNodes()
					sol, err := solve.FromNodesWithTimeout(
						iter.GetWidth(),
						iter.GetHeight(),
						ns,
						time.Duration(iter+1)*100*time.Millisecond,
					)
					if err != nil {
						t.Logf("Error fetching input: %q", err)
						t.Fail()
					}
					if sol.ToAnswer() != sr.Answer {
						t.Logf("Incorrect Answer\n")
						t.Logf("Exp: %q\n", sr.Answer)
						t.Logf("Act: %q\n", sol.ToAnswer())
						t.Logf("Board:\n%s\n\n", sol.Pretty(ns))
						t.Fail()
					}
				})
			}
		})
	}
}

func BenchmarkAll(b *testing.B) {
	fetch.DisableHTTPCalls()
	solve.SetTestTimeout()
	// go decided that it should run benchmarks in this directory.
	os.Chdir(`..`)

	for iter := model.MinIterator; iter <= model.MaxIterator; iter++ {
		b.Run(iter.String(), func(b *testing.B) {
			srs, err := fetch.ReadN(iter, 10)
			if err != nil {
				b.Logf("Error fetching input: %q", err)
				b.Fail()
			} else if len(srs) == 0 {
				b.Logf("No cached results")
				b.Fail()
			}

			for _, sr := range srs {
				sr := sr
				if sr.Answer == `` {
					b.Logf("Unknown answer: %+v", sr)
					b.Fail()
				}
				b.Run("PuzzleID "+sr.Input.ID, func(b *testing.B) {
					var sol model.Solution
					for n := 0; n < b.N; n++ {
						sol, err = solve.FromNodesWithTimeout(
							iter.GetWidth(),
							iter.GetHeight(),
							sr.Input.ToNodes(),
							time.Duration(iter+1)*100*time.Millisecond,
						)
						if err != nil {
							b.Logf("got unexpected error: %q", err)
							b.Fail()
						}
						if sol.ToAnswer() != sr.Answer {
							b.Logf("expected answer %q, but got %q", sr.Answer, sol.ToAnswer())
							b.Fail()
						}
					}
				})
			}
		})
	}
}
*/
