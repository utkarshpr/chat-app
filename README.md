Here is the `README.md` file formatted in a way you can directly copy and paste:

```markdown
# Real-Time Chat Application API Documentation

This is the API documentation for the Real-Time Chat Application, designed to facilitate user authentication, message handling, contact management, and user profile updates. The API follows the Swagger 2.0 specification.

## Getting Started

To run the API and view the documentation, follow the instructions below.

### Prerequisites

1. **Go** - Version 1.18 or higher
2. **MongoDB** - For storing user data and messages
3. **Swagger UI** - To visualize and interact with the API endpoints

### Running the Application Locally

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-repo/real-time-chat-app.git
   cd real-time-chat-app
   ```

2. **Install dependencies**:

   If you have Go modules enabled, you can install the necessary dependencies by running:

   ```bash
   go mod tidy
   ```

3. **Run the server**:

   To start the application, run:

   ```bash
   go run main.go
   ```

4. **Access the API**:

   The server will start, and you can access the API on [http://localhost:8080](http://localhost:8080).

### Swagger UI

To view the API documentation, navigate to:

- [Swagger UI for the API Documentation](http://localhost:8080/swagger/index.html)

You can interact with the API through the Swagger UI, which is generated based on the `swagger.yaml` file located in the `docs` folder.

Alternatively, you can access the raw Swagger JSON directly by visiting:

- [Swagger JSON](http://localhost:8080/swagger.json)

### API Endpoints

#### 1. Authentication

- **POST /auth/login**  
  Authenticates the user with valid credentials and returns a JWT token.

- **POST /auth/logout**  
  Logs out the currently logged-in user by invalidating their JWT token.

- **POST /auth/signup**  
  Processes user registration by validating the input and creating a new user.

#### 2. Contacts

- **POST /contacts/action**  
  Handles the blocking or removal of a contact based on user input.

- **POST /contacts/add**  
  Adds or updates a contact request between users.

#### 3. Messages

- **GET /messages/get**  
  Fetches all messages exchanged with a specific recipient.

- **POST /messages/sent**  
  Sends a new message from the authorized user to the recipient.

#### 4. User Management

- **DELETE /user/deleteUser**  
  Deletes a user account by the specified username.

- **GET /user/fetchUser**  
  Retrieves user details using a username provided in the query parameters.

- **POST /user/updateUserAndProfile**  
  Updates the details of a user based on the provided username and JSON body payload.

### API Definitions

- **models.GenericResponse**  
  A generic response structure used across all endpoints.

- **models.LoginUser**  
  Contains the `username` and `password` fields for user login.

- **models.User**  
  Contains user details like `first_name`, `last_name`, `email`, etc.

- **models.Message**  
  Contains message details such as `chat_id`, `sender_id`, `recipient_id`, etc.

- **models.ContactRequest**  
  Used to add a contact request between users.

- **models.ContactActionRequest**  
  Defines actions like `remove` or `block` for managing contacts.

### File Structure

- `main.go` - Entry point of the application, sets up routes and API handlers.
- `docs/swagger.yaml` - Contains the Swagger API documentation in YAML format.
- `routes/` - Contains route definitions for user authentication, contacts, and messaging.
- `models/` - Contains all data models used in the API.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- **Swagger**: For generating the API documentation.
- **Gin**: A web framework for Go used to build the API.
- **MongoDB**: For the database used in storing user and message data.

---

Feel free to open issues for any questions, bugs, or suggestions!
```

### Notes:
- The provided markdown ensures that you can copy the entire content directly into your README file without any formatting issues.
- Make sure to replace `https://github.com/your-repo/real-time-chat-app.git` with the actual URL of your repository.
- The structure includes detailed information about the API, endpoints, and available models. Adjust as necessary based on your specific API implementation.
