$(document).ready(function() {
    $('#loginForm').on('submit', function(e) {
        e.preventDefault();

        var userData = {
            username: $('#loginForm input[name="username"]').val(),
            password: $('#loginForm input[name="password"]').val()
        };

        $.ajax({
            type: "POST",
            url: "/login",
            contentType: "application/json",
            data: JSON.stringify(userData),
            success: function(response) {
                console.log(response);
                localStorage.setItem("token", response.token); // Store the token in localStorage
                window.location.href = '/some-protected-route'; // Redirect to the protected page
            },
            error: function(xhr, status, error) {
                var response = JSON.parse(xhr.responseText);
                console.log(xhr.responseText);
                alert("Login failed: " + response.error);
            }
        });
    });
});
