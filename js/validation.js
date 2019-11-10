// returns error message
function validate(element) {
    var classes = element.classList
    var errorElement = element.nextElementSibling
    var allErrors = ""

    if (isIn(classes, 'form-required')) {
        if (element.value < 1) {
            var errorString = "Pole " + element.getAttribute("name") + " jest wymagane"
            addErrorMessage(errorString, errorElement)
            allErrors += errorString
        }
    }

    // if passed all checks
    if (allErrors == "") {
        clearErrorMessage(errorElement)
    }

    return allErrors
}

function validateForm(element) {
    // validate all form elements
    // display combined error message
    var children = element.childNodes
    var errorBox = document.getElementById("error-box")
    var allErrors = []

    for (var i=0; i<children.length; i++) {
        if(children[i].classList != undefined && isIn(children[i].classList, 'validate')) {
            var validationResult = validate(children[i])
            if (validationResult != "") {
                allErrors.push(validationResult)
            }
        }
    }

    console.log(allErrors)
    if (allErrors.length == 0) {
        return true
    } else {
        errorBox.innerHTML = ""
        for (var i = 0; i < allErrors.length; i++) {
            errorBox.innerHTML += allErrors[i] + "<br>"
        }
        return false
    }
}

function addErrorMessage(msg, errorElement) {
    if (!isIn(errorElement.classList, 'error-message')) {
        console.log('error element does not contain error-message class!')
    }
    errorElement.innerHTML = msg
}

function clearErrorMessage(errorElement) {
    if (!isIn(errorElement.classList, 'error-message')) {
        console.log('error element does not contain error-message class!')
    }
    errorElement.innerHTML = ''
}

function isIn(array, element) {
    for (var i=0; i<array.length; i++) {
        if (element == array[i]) {
            return true
        }
    }
    return false
}