// assets/js/index.js

$(document).ready(function() {
    $.ajax({
        url: '/user-manager/api/users',
        type: 'GET',
        dataType: 'json',
        success: function(response) {
            var tbodyContent = '';
            response.users.forEach(function(user) {
                var userRow = `
                    <tr class="cursor-pointer user-row bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600" data-user-id="${user.ID}">
                        <td class="px-6 py-4 text-gray-900 whitespace-nowrap dark:text-white">
                            ${user.ID}
                        </td>
                        <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            ${user.username}
                        </td>
                        <td class="px-6 py-4">
                            ${new Date(user.CreatedAt).toLocaleDateString()} ${new Date(user.CreatedAt).toLocaleTimeString()}
                        </td>
                        <td class="px-6 py-4 text-right">
                            <button id="edit-${user.ID}" class="edit-btn font-medium text-blue-600 dark:text-blue-500 hover:underline mr-3">Edit</button>
                            <button id="delete-${user.ID}" class="delete-btn font-medium text-red-600 dark:text-red-500 hover:underline">Delete</button>
                        </td>
                    </tr>
                `;
                tbodyContent += userRow;
            });
            // Use the ID to target the table body and update its content
            $('#usersTableBody').html(tbodyContent);

            // Add click event listener for each user row
            $('.user-row').click(function(event) {
                // Prevent redirection when clicking on Edit or Delete buttons
                if (!$(event.target).hasClass('edit-btn') && !$(event.target).hasClass('delete-btn')) {
                    var userID = $(this).data('user-id');
                    window.location.href = `/user-manager/users/${userID}`;
                }
            });
        },
        error: function(error) {
            var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
            var errorAlert = $('#error-message');
            errorAlert.text(errorMsg);
            errorAlert.removeClass('hidden');
            console.error('Error fetching users:', error);
        }
    });
});
