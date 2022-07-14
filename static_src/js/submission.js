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
        var copyFrom = document.createElement("textarea");
        copyFrom.textContent = textVal;
        var bodyElm = document.getElementsByTagName("body")[0];
        bodyElm.appendChild(copyFrom);
        copyFrom.select();
        var retVal = document.execCommand('copy');
        bodyElm.removeChild(copyFrom);
        return retVal;
    }
}

$(function() {
    $("#source-expand").click(function() {
        const code = $('#'+$(this).data('target'));
        if(code.data('expanded')) {
            code.data('expanded', false);
            code.css('max-height', '400px');
        }else {
            code.data('expanded', true);
            code.css('max-height', 'none');
        }
    });

    $("#source-copy").click(function() {
       window.getSelection().removeAllRanges();
       copy($('#'+$(this).data('target')).text());
    });
})