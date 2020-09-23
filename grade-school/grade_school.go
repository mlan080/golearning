package school

import "sort"

type Grade struct {
	grade    int
	students []string
}

type School struct {
	grades map[int][]string
}

func New() *School {
	return &School{
		grades: make(map[int][]string, 0),
	}
}

//adds student into the grade
func (s *School) Add(student string, grade int) {
	s.grades[grade] = append(s.grades[grade], student)
}

// //returns students for given grade
// - Get a list of all students enrolled in a grade
func (s *School) Grade(grade int) []string {
	return s.grades[grade]
}

//  //Get a sorted list of all students in all grades. Grades should sort as 1, 2, 3, etc., and students within a grade should be sorted alphabetically by name.
// //"Who all is enrolled in school right now?"
// //"Grade 1: Anna, Barb, and Charlie. Grade 2: Alex, Peter, and Zoe. Grade 3…"
func (s *School) Enrollment() []Grade {
	result := make([]Grade, len(s.grades))

	keys := make([]int, len(s.grades))
	i := 0
	for k := range s.grades {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	for i, k := range keys {
		v := s.grades[k]
		sort.Strings(v)
		result[i] = Grade{
			grade:    k,
			students: v,
		}
	}
	return result
}
