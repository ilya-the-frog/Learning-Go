//проверяем, все ли скобки в строке открыты и закрыты корректно

func check(str byte[]) bool {
    for i := 0; i < len(str); i++ {
        if str[i] == "[" || str[i] == "("
            { stack.add(str[i]) }

        if str[i] == "]" {
            if "[" == stack.pop()
                { continue }
             else
                { return false }
        }

        if str[i] == ")" {
            if "(" == stack.pop()
                { continue }
            else
                { return false }

        }
    }

    if stack == nil {
            return true
        } else {
            return false
        }
}
