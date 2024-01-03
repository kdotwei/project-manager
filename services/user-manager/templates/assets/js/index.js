// assets/js/index.js

$(document).ready(function() {
    $('#createButton').click(function() {
        window.location.href = '/user-manager/users/create';
    });

    // Load users
    loadUsers();

    // Function to load users
    function loadUsers() {
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
                $('#usersTableBody').html(tbodyContent);

                // Click event listener for Edit buttons
                $('.edit-btn').click(function(event) {
                    event.stopPropagation(); // Prevent triggering the user-row click event
                    var user_id = $(this).attr('id').split('-')[1];
                    window.location.href = `/user-manager/users/${user_id}/edit`;
                });

                // Click event listener for Delete buttons
                $('.delete-btn').click(function(event) {
                    event.stopPropagation(); // Prevent triggering the user-row click event
                    var user_id = $(this).attr('id').split('-')[1];
                    deleteUser(user_id);
                });

                // Click event listener for user row
                $('.user-row').click(function(event) {
                    if (!$(event.target).hasClass('edit-btn') && !$(event.target).hasClass('delete-btn')) {
                        var user_id = $(this).data('user-id');
                        window.location.href = `/user-manager/users/${user_id}`;
                    }
                });
            },
            error: function(error) {
                // Error handling
            }
        });
    }

    // Function to delete user
    function deleteUser(user_id) {
        if (confirm('Are you sure you want to delete this user?')) {
            $.ajax({
                url: `/user-manager/api/users/${user_id}/delete`,
                type: 'DELETE',
                success: function(response) {
                    console.log('User deleted:', response);
                    $('#success-message').text(response.message).removeClass('hidden');
                    loadUsers(); // Reload users after deletion
                },
                error: function(error) {
                    var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                    $('#error-message').text(errorMsg).removeClass('hidden');
                    console.error('Error deleting user:', error);
                }
            });
        }
    }
});
