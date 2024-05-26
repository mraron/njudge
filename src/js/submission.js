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

    $("#source-copy").click(copyHandler);
})