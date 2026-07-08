func generateParenthesis(n int) []string {
   res := make([]string, 0)

   var valid func(string) bool
   valid = func(s string) bool {
       open := 0
       for _, c := range s {
           if c == '(' {
               open++
           } else {
               open--
           }
           if open < 0 {
               return false
           }
       }
       return open == 0
   }

   var dfs func(string)
   dfs = func(s string) {
       if len(s) == n*2 {
           if valid(s) {
               res = append(res, s)
           }
           return
       }

       dfs(s + "(")
       dfs(s + ")")
   }

   dfs("")
   return res
}