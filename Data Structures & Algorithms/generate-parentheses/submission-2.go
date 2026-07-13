func generateParenthesis(n int) []string {
   stack := make([]string, 0)
   res := make([]string, 0)

   var backtrack func(int, int)
   backtrack = func(openN, closedN int) {
       if openN == n && closedN == n {
           res = append(res, strings.Join(stack, ""))
           return
       }

       if openN < n {
           stack = append(stack, "(")
           backtrack(openN+1, closedN)
           stack = stack[:len(stack)-1]
       }

       if closedN < openN {
           stack = append(stack, ")")
           backtrack(openN, closedN+1)
           stack = stack[:len(stack)-1]
       }
   }

   backtrack(0, 0)
   return res
}