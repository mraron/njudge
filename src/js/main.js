function copy(textVal) {
    if (navigator.clipboard) {
        let result;
        navigator.clipboard.writeText(textVal).then(() => {
            result = true;
        }).catch(err => {
            result = false;
            console.log(err)
        });

        return result;
    }else {
        let copyFrom = document.createElement("textarea");
        copyFrom.textContent = textVal;
        let bodyElm = document.getElementsByTagName("body")[0];
        bodyElm.appendChild(copyFrom);
        copyFrom.select();
        let retVal = document.execCommand('copy');
        bodyElm.removeChild(copyFrom);
        return retVal;
    }
}

function copyHandler() {
    window.getSelection().removeAllRanges();
    copy($('#'+$(this).data('target')).text());
}

$(".input-output-copier").click(copyHandler);