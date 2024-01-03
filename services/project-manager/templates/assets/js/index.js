// assets/js/index.js

$(document).ready(function() {
    $('#createProjectForm').submit(function(event) {
        event.preventDefault();
        createProject()
    });

    // Load projects
    loadProjects();

    // Function to load projects
    function loadProjects() {
        $.ajax({
            url: '/project-manager/api/projects',
            type: 'GET',
            dataType: 'json',
            success: function(response) {
                console.log(response)
                var tbodyContent = '';
                response.projects.forEach(function(project) {
                    var projectRow = `
                        <tr class="cursor-pointer project-row bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600" data-project-id="${project.ID}">
                            <td class="px-6 py-4 text-gray-900 whitespace-nowrap dark:text-white">${project.ID}</td>
                            <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">${project.name}</td>
                            <td class="px-6 py-4">${new Date(project.CreatedAt).toLocaleDateString()} ${new Date(project.CreatedAt).toLocaleTimeString()}</td>
                            <td class="px-6 py-4">${new Date(project.UpdatedAt).toLocaleDateString()} ${new Date(project.UpdatedAt).toLocaleTimeString()}</td>
                            <td class="px-6 py-4 text-right">
                                <button id="edit-${project.ID}" class="edit-btn font-medium text-blue-600 dark:text-blue-500 hover:underline mr-3">Edit</button>
                                <button id="delete-${project.ID}" class="delete-btn font-medium text-red-600 dark:text-red-500 hover:underline">Delete</button>
                            </td>
                        </tr>
                    `;
                    tbodyContent += projectRow;
                });
                $('#projectsTableBody').html(tbodyContent);

                // Click event listener for Edit buttons
                $('.edit-btn').click(function(event) {
                    event.stopPropagation(); // Prevent triggering the project-row click event
                    var projectId = $(this).attr('id').split('-')[1];
                    window.location.href = `/project-manager/projects/${projectId}/edit`;
                });
                
                // Click event listener for Delete buttons
                $('.delete-btn').click(function(event) {
                    event.stopPropagation(); // Prevent triggering the project-row click event
                    var projectId = $(this).attr('id').split('-')[1];
                    deleteProject(projectId);
                });

                // Click event listener for project row
                $('.project-row').click(function(event) {
                    if (!$(event.target).hasClass('edit-btn') && !$(event.target).hasClass('delete-btn')) {
                        var projectId = $(this).data('project-id');
                        window.location.href = `/project-manager/projects/${projectId}/tasks`;
                    }
                });
            },
            error: function(error) {
                // Error handling
            }
        });
    }

    // Function to delete project
    function deleteProject(projectId) {
        if (confirm('Are you sure you want to delete this project?')) {
            $.ajax({
                url: `/project-manager/api/projects/${projectId}/delete`,
                type: 'DELETE',
                success: function(response) {
                    console.log('Project deleted:', response);
                    $('#success-message').text(response.message).removeClass('hidden');
                    loadProjects(); // Reload projects after deletion
                },
                error: function(error) {
                    var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                    $('#error-message').text(error_message).removeClass('hidden');
                    console.error('Error deleting project:', error);
                }
            });
        }
    }

    function createProject() {
        var project_name = $('#project-name').val();
        var data = {
            Name: project_name
        }

        $.ajax({
            type: 'POST',
            url: '/project-manager/api/projects/create',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function(response) {
                console.log('Project created successfully:', response);
                $('#success-message').text(response.message).removeClass('hidden');
                loadProjects(); // Reload projects after creation
            },
            error: function(error) {
                var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                $('#error-message').text(error_message).removeClass('hidden');
                console.error('Error fetching users:', error);
            }
        });
    }
});
