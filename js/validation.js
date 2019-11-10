function validate(element) {
    var classes = element.classList
    var errorElement = element.nextElementSibling
    var ok = true

    if (isIn(classes, 'form-required')) {
        if (element.value < 1) {
            addErrorMessage("This field is required", errorElement)
            ok = false
        }
    }
    // if passed all checks
    if (ok) {
        clearErrorMessage(errorElement)
    }

    return ok
}

function validateForm(element) {
    // validate all form elements
    // display combined error message
    var children = element.getChildren
    for (var i=0; i<children.length; i++) {
        if(isIn(children[i].classList, 'validate')) {
            validate(children[i])
        }
    }
}

function addErrorMessage(msg, errorElement) {
    if (!isIn(errorElement.classList, 'error-message')) {
        console.log('error element does not contain error-message class!')
    }
    errorElement.innerHTML = 'To pole jest wymagane'
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