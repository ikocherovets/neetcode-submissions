func checkInclusion(s1 string, s2 string) bool {
    // If s1 is longer, a permutation can't fit. Immediate no.
    if len(s1) > len(s2) {
        return false
    }

    // Two passports of 26 slots each.
    // need   — the reference (s1), built once, never changes.
    // window — the current window's passport, changes as it slides.
    var need, window [26]int

    // Build s1's passport and the first window (first len(s1) chars of s2).
    for i := 0; i < len(s1); i++ {
        need[s1[i]-'a']++
        window[s2[i]-'a']++
    }

    // Maybe the very first window matches.
    if need == window {
        return true
    }

    // Roll the window across s2, left to right.
    for r := len(s1); r < len(s2); r++ {
        window[s2[r]-'a']++         // letter entering on the right (position r)
        window[s2[r-len(s1)]-'a']-- // letter leaving on the left (position r-len(s1))
        if need == window {         // passports match?
            return true
        }
    }

    // Reached the end with no match.
    return false
}