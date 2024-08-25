package picine

func RectPerimeter(w, h int) int {
    if w == 0 || h == 0{
        return 0
    }
    if w < 0 || h < 0{
        a := -1
        return a
    }else{
        f := w + h
        ans := f * 2
        return ans

    }
}