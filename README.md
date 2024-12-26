Here’s the updated `README.md` without the exact values of the `.env` file:

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

3. **Configure the environment**:

   Create a `.env` file in the root directory of the project and add the following configuration:

   ```env
   MONGO_URI=<your-mongo-uri>
   MONGO_DATABASE=<your-database-name>
   MONGO_TABLE_USER=<your-user-table>
   MONGO_TABLE_JWT_STORE=<your-jwt-table>
   MONGO_TABLE_CONTACT=<your-contact-table>
   MONGO_TABLE_MESSAGE=<your-message-table>

   PORT=:8081

   JWT_SECRET_KEY=<your-jwt-secret-key>
   ```

   Replace the placeholders with the appropriate values for your setup.

4. **Run the server**:

   To start the application, run:

   ```bash
   go run main.go
   ```

5. **Access the API**:

   The server will start, and you can access the API on [http://localhost:8081](http://localhost:8081).

### Swagger UI

To view the API documentation, navigate to:

- [Swagger UI for the API Documentation](http://localhost:8081/swagger/index.html)

You can interact with the API through the Swagger UI, which is generated based on the `swagger.yaml` file located in the `docs` folder.

Alternatively, you can access the raw Swagger JSON directly by visiting:

- [Swagger JSON](http://localhost:8081/swagger.json)

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

### Environment Variables

Below are the environment variables used by the application:

- `MONGO_URI`: The MongoDB connection URI.
- `MONGO_DATABASE`: The name of the MongoDB database.
- `MONGO_TABLE_USER`: The name of the user table in the database.
- `MONGO_TABLE_JWT_STORE`: The table to store JWT tokens.
- `MONGO_TABLE_CONTACT`: The table to store contact information.
- `MONGO_TABLE_MESSAGE`: The table to store messages.
- `PORT`: The port number for the application to listen on.
- `JWT_SECRET_KEY`: The secret key used for signing JWT tokens.

Make sure to replace the placeholders in the `.env` file with your actual values.
```

This version hides the sensitive values in the `.env` file while keeping the structure intact. You can copy and paste this into your `README.md` directly.
