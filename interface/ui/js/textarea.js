
const textArea = document.getElementById("area-helper")



textArea.addEventListener("keydown",function(e) {

    const pairs = {
        "{":"{}",
        "[":"[]",
        '"':'""'
    }

    let initialPos = textArea.selectionStart
    let beforText = textArea.value.substring(0,initialPos)
    let afterText = textArea.value.substring(initialPos)
    let offsetLen = 0
    let borderOffset = ""

    if (pairs[e.key]) {
        e.preventDefault()
        insertText = pairs[e.key]
        offsetLen = 1
    }
    if (e.key === "Enter") {
        e.preventDefault()
        currentLine = beforText.split("\n").at(-1)
        borderOffset = currentLine.includes('"') ? currentLine.split('"')[0] : ""
        if ((afterText[0] === "]") || (afterText[0] === "}")) {
            offsetLen = borderOffset.length + 3
            insertText = "\n" + " ".repeat(offsetLen) + "\n" + borderOffset
        } else if (beforText.at(-1) === ",") {
            offsetLen = borderOffset.length + 1
            insertText = "\n" + " ".repeat(offsetLen)
        }
    }

    if (offsetLen) {
        textArea.value = beforText + insertText + afterText
        textArea.selectionStart = initialPos + offsetLen
        textArea.selectionEnd = initialPos + offsetLen
    }



    
})