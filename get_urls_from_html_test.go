package main

import "testing"

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no links",
			inputURL: "https://example.com",
			inputBody: `
		<html>
			<body>
				<p>No links here</p>
			</body>
		</html>
		`,
			expected: []string{},
		},
		{
			name:     "links with fragments",
			inputURL: "https://example.com/base",
			inputBody: `
		<html>
			<body>
				<a href="/path#section">Section Link</a>
				<a href="https://example.com/another#part">Another Link</a>
			</body>
		</html>
		`,
			expected: []string{"https://example.com/path#section", "https://example.com/another#part"},
		},
		{
			name:     "deeper nested relative url",
			inputURL: "https://example.com/some/subdir/index.html",
			inputBody: `
			<html>
				<body>
					<div>
						<section>
							<div>
								<a href="../../another/subdir">Deep Link</a>
							</div>
						</section>
					</div>
				</body>
			</html>
			`,
			expected: []string{"https://example.com/another/subdir"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: expected no error but got one", i, tc.name)
			}

			for i, s := range actual {
				if s != actual[i] {
					t.Errorf("Test %v - '%s' FAIL: expected %v but got %v", i, tc.name, tc.expected, actual)
				}
			}

		})
	}
}
