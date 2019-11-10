// returns error message
function validate(element) {
    var classes = element.classList
    var errorElement = element.nextElementSibling
    var allErrors = ""

    clearErrorMessage(errorElement)
    element.style.border = "2px solid forestgreen"

    if (isIn(classes, 'form-required') && element.value < 1) {
        var errorString = "Pole " + element.getAttribute("name") + " jest wymagane "
        addErrorMessage(errorString, errorElement)
        allErrors += errorString
    }
    if (isIn(classes, 'form-email') && !validateEmail(element.value)) {
        var errorString = "Niepoprawny adres email "
        addErrorMessage(errorString, errorElement)
        allErrors += errorString
    }

    // add invalid field styling
    if (allErrors != "") {
        element.style.border = "2px solid crimson"
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

function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}

function addErrorMessage(msg, errorElement) {
    if (!isIn(errorElement.classList, 'error-message')) {
        console.log('error element does not contain error-message class!')
    }
    if (errorElement.innerHTML != "") {
        errorElement.innerHTML += "<br>"
    }
    errorElement.innerHTML += msg
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