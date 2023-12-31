$(document).ready(function () {
    $('#registerForm').on('submit', function (e) {
        e.preventDefault();

        var userData = {
            username: $('#username').val(),
            password: $('#password').val()
        };

        $.ajax({
            type: "POST",
            url: "/identifier/register",
            contentType: "application/json",
            data: JSON.stringify(userData),
            success: function (response) {
                console.log(response);
                window.location.href = '/identifier/login'; // Redirect to the protected page
            },
            error: function (xhr, status, error) {
                var response = JSON.parse(xhr.responseText);
                console.log(xhr.responseText);
                $('#registerError').html("Register failed: " + response.error).addClass('error-msg');
            }
        });
    });
});