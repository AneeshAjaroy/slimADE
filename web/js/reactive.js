
const textArea = document.getElementById("area-helper")
const selectArea = document.getElementById("select-helper")
const queryAdd = document.getElementById("query-add-button")
const headerAdd = document.getElementById("header-add-button")




function UpdateColor() {
    const color = selectArea.options[selectArea.selectedIndex].style.color;
    console.log("hi",color)
    selectArea.style.backgroundColor = color;
}

UpdateColor()

selectArea.addEventListener("change",UpdateColor)

textArea.addEventListener("keydown",function(e) {

    const pairs = {
        "{":"}",
        "[":"]",
        '"':'"'
    }

    const pairsCom = {
        "{":"{}",
        "[":"[]",
        '"':'""'
    }

    const revPairs = {
        "}":"{",
        "]":"[",
        '"':'"'
    }

    let initialPos = textArea.selectionStart
    let finalPos = textArea.selectionEnd
    let beforeText = textArea.value.substring(0,initialPos)
    let afterText = textArea.value.substring(finalPos)
    let offsetLen = 0
    let borderOffset = ""
    let insertText = ""


    if (pairs[e.key]) {
        if (finalPos !== initialPos) {
            e.preventDefault()
            console.log(textArea.value.substring(initialPos,finalPos))
            insertText = e.key + textArea.value.substring(initialPos,finalPos) + pairs[e.key]
            offsetLen = finalPos - initialPos + 3
        } else {
            e.preventDefault()
            insertText = e.key + pairs[e.key]
            offsetLen = 1
        }

    }

    if (revPairs[e.key] && afterText[0] === e.key) {
        e.preventDefault()
        offsetLen = 1
    }

    if (e.key === "Backspace") {
        beforeKey = beforeText.at(-1)
        afterKey = afterText[0]
        if (pairsCom[beforeKey] === beforeKey + afterKey) {
            e.preventDefault()
            offsetLen = -1
            beforeText = beforeText.slice(0,-1)
            afterText = afterText.slice(1)
        }
    }

    if (e.key === "Enter") {
        e.preventDefault()
        currentLine = beforeText.split("\n").at(-1)
        borderOffset = currentLine.includes('"') ? currentLine.split('"')[0] : ""
        if ((afterText[0] === "]") || (afterText[0] === "}")) {
            offsetLen = borderOffset.length + 3
            insertText = "\n" + " ".repeat(offsetLen) + "\n" + borderOffset
        } else if (beforeText.at(-1) === ",") {
            offsetLen = borderOffset.length + 1
            insertText = "\n" + " ".repeat(offsetLen)
        }
    }

    if (offsetLen !== 0) {
        textArea.value = beforeText + insertText + afterText
        textArea.selectionStart = initialPos + offsetLen
        textArea.selectionEnd = initialPos + offsetLen
    }

})


queryAdd.addEventListener("click",function(e) {
    tableBody = document.querySelector("#queryTable tbody")
    tableBody.insertAdjacentHTML("beforeend", `
    <tr>
        <td >
            <input name="enabled-q" class="form-check-input border-black" type="hidden"  value="false" style="width: 100%;">
            <input name="enabled-q" class="form-check-input border-black" type="checkbox" value="true"  style="width: 100%;">
        </td>
        <td>
            <input name="key-q" type="text" class="form-control border-black" style="overflow: scroll;">
        </td>
        <td>
            <input name="value-q" type="text" class="form-control border-black" style="overflow: scroll;">
        </td>
        <td>
            <button type="button" class="btn-close query-remove" style="width: 100%; box-sizing:border-box"></button>
        </td>
    </tr>
`);
    deleteButtons = document.querySelectorAll(".query-remove");

    deleteButtons.forEach(button => {
    button.addEventListener("click", function() {
            const btnList = document.querySelectorAll(".query-remove");
            if (btnList.length !== 1) {
                const row = this.closest("tr"); 
                row.remove();
            }

        });
    })
})



headerAdd.addEventListener("click",function(e) {
    tableBody = document.querySelector("#headerTable tbody")
    tableBody.insertAdjacentHTML("beforeend", `
    <tr>
        <td >
            <input name="enabled-h" class="form-check-input border-black" type="hidden" value="false"  style="width: 100%;">
            <input name="enabled-h" class="form-check-input border-black" type="checkbox" value="true" style="width: 100%;">
        </td>
        <td>
            <input name="key-h" type="text" class="form-control border-black" style="overflow: scroll;">
        </td>
        <td>
            <input name="value-h"  type="text" class="form-control border-black" style="overflow: scroll;">
        </td>
        <td>
            <button type="button" class="btn-close header-remove" style="width: 100%; box-sizing:border-box"></button>
        </td>
    </tr>
`);
    deleteButtons = document.querySelectorAll(".header-remove");

    deleteButtons.forEach(button => {
    button.addEventListener("click", function() {
            const btnList = document.querySelectorAll(".header-remove");
            if (btnList.length !== 1) {
                const row = this.closest("tr"); 
                row.remove();
            }

        });
    })
})



