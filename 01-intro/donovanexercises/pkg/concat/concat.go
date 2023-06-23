/* String concatenation
Exercise 1.3: Experiment to measure the difference in running time between our potentially
inefficient versions and the one that uses strings.Join.

Плюс я попробовал сделать функцию быстрее чем strings.Join.
Получилось так себе, лучше использовать strings.Join :D
*/

package concat

import (
	"fmt"
)

func basic(ss []string, s *string) error {
	switch len(ss) {
	case 0:
		return fmt.Errorf("slice should contain at least one element")
	case 1:
		*s = ss[0]
		return nil
	}

	sep := " "
	*s = ss[0]
	for _, v := range ss[1:] {
		*s += sep + v
	}

	return nil
}

func custom(ss []string, sep string) string {
	switch len(ss) {
	case 0:
		return ""
	case 1:
		return ss[0]
	}

	n := len(sep) * (len(ss) - 1)
	for i := 0; i < len(ss); i++ {
		n += len(ss[i])
	}
	s := make([]byte, 0, n)
	s = append(s, ss[0]...)
	for _, v := range ss[1:] {
		s = append(s, sep...)
		s = append(s, v...)
	}
	return string(s)
}
