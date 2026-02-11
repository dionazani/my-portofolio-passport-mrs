# my-portofolio-passport-mrs

## Background
System for a movie reservation service as describe in roadmap.sh project: https://roadmap.sh/projects/movie-reservation-system

## Goal
The goal of this project is to help you understand how to implement complex business logic i.e. seat reservation and scheduling, thinking about the data model and relationships, and complex queries.

## Requirements
We have intentionally left out the implementation details to encourage you to think about the design and implementation of the system. However here are some requirements that you can consider:

## Main Feature

#### User Authentication and Authorization
1. Users should be able to sign up and log in.
2. You also need roles for users, such as admin and regular user. Admins should be able to manage movies and showtimes.
3. Regular users should be able to reserve seats for a showtime.

### Movie Management
Admins should be able to add, update, and delete movies.

1. Each movie should have a title, description, and poster image.
2. Movies should be categorized by genre.
3. Movies should have showtimes.

### Reservation Management

1. Think about the data model and relationships between entities.
2. Think about how you will avoid overbooking, and how you will handle seat reservations.
3. Think about how you will handle the scheduling of showtimes.
4. Think about how you will handle the reporting of reservations.
5. Think about how you will handle the authentication and authorization of users.

### ToDo List

| No. | Feature | Status | Remarks |
| :--- | :--- | :--- | :--- |
| 1. | Think about the data model and relationships between entities. | **Done** | Please see folder: `database` |
| 2. | Users should be able to sign up and log in | **Done** | Please see folder: `\backend-go\sign-up-service` and `\backup-go\auth-service` |
| 3. | You also need roles for users, such as admin and regular user. Admins should be able to manage movies and showtimes. | **In Progress** | - |

### API Collection
https://web.postman.co/workspace/My-Workspace~69337442-700a-4ec7-85eb-9ca980d54a45/collection/5032744-9e225fb2-2831-46ce-b2f3-b6197deb8a32?action=share&source=copy-link&creator=5032744
