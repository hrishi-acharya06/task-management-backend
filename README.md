# task-management-backend
CREATE USER 'task_management_user'@'localhost' IDENTIFIED BY 'task_management';
GRANT ALL PRIVILEGES ON task_management.* TO 'task_management_user'@'localhost';