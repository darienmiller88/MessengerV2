package sqlconstants

const (
	///////////////////////////////////////////
	// CREATE
	///////////////////////////////////////////

	// Insert a message into "messages" with a particular chat_id.
	INSERT_GROUP_CHAT_MESSAGE string = "INSERT INTO messages " +
	"(message_content, message_date, created_at, updated_at, username, image_url, display_name, chat_id) " +
	"VALUES(:message_content, :message_date, :created_at, :updated_at, :username, :image_url, :display_name, :chat_id) " +
	"RETURNING id"

	// Insert a public message into "messages".
	INSERT_PUBLIC_MESSAGE string = "INSERT INTO messages " +
	"(message_content, message_date, created_at, updated_at, username, image_url, display_name) " +
	"VALUES(:message_content, :message_date, :created_at, :updated_at, :username, :image_url, :display_name) " +
	"RETURNING id"

	// Insert a new user into the "users" table.
	INSERT_USER string = "INSERT INTO users (created_at, updated_at, username, password, is_anonymous) " +
	"VALUES (:created_at, :updated_at, :username, :password, :is_anonymous) RETURNING id"

	//Insert a new group chat into the "chats" table.
	INSERT_GROUP_CHAT string = "INSERT INTO chats (created_at, updated_at, chat_name, picture_url) " +
	"VALUES(:created_at, :updated_at, :chat_name, :picture_url) RETURNING id"

	// Insert a new "user_chat" relation into the "user_chats" table.
	INSERT_USER_CHAT string = "INSERT INTO user_chats (created_at, updated_at, chat_id, username) " +
	"VALUES(:created_at, :updated_at, :chat_id, :username) RETURNING id"



	////////////////////////////////////////////////////////
	// READ
	///////////////////////////////////////////////////////

	//SQL statement that joins the messages and users table in order to retrieve all messages, and the profile
	// picture of the user that sent each message.
	GET_PUBLIC_MESSAGES string = "SELECT messages.*, users.profile_picture FROM messages JOIN users ON " + 
	"messages.username = users.username WHERE chat_id IS NULL ORDER BY id ASC"

	//Retrive the same information as public messages, but only for private group chat messages.
	GET_GROUP_CHAT_MESSAGES string = "SELECT messages.*, users.profile_picture FROM messages JOIN users ON " +
	"messages.username = users.username WHERE chat_id=$1 ORDER BY id ASC"

	//Get the messages belonging to a specific user.
	GET_MESSAGE_HISTORY string = "SELECT * FROM messages WHERE username=$1 ORDER BY id ASC"

	//Get all of the group chats belonging to a particular user by joining "chats" onto "user_chats", and
	//then by joining 'user_chats" onto "users". 
	GET_GROUP_CHATS string = "SELECT chats.* FROM chats JOIN user_chats ON chats.id = user_chats.chat_id " +
	"JOIN users ON users.username = user_chats.username WHERE users.username=$1"

	//Get all users in the "users" table, only retrieving the username, profile_picture, and is_anonymous columns.
	GET_ALL_USERS string = "SELECT username, profile_picture, is_anonymous FROM users"

	//Get a user by their username.
	GET_USER_BY_USERNAME string = "SELECT * FROM users WHERE username=$1"




	////////////////////////////////////////////////
	//UPDATE
	/////////////////////////////////////////////////
	
	//Update the user's profile picture and display name
	UPDATE_USER_PROFILE_PICTURE string = "UPDATE users SET display_name=$1, profile_picture=$2, updated_at=$3 WHERE username=$4"

	//Update the user's display name only
	UPDATE_USER string = "UPDATE users SET display_name=$1, updated_at=$2 WHERE username=$3"



	
	///////////////////////////////////////////////////
	// DESTROY
	///////////////////////////////////////////////////

	//Delete a message from "messages" with a particular id and username.
	DELETE_MESSAGE string = "DELETE FROM messages WHERE id=$1 and username=$2"

	//Delete a group chat from "chats" with a particular id.
	DELETE_CHAT string = "DELETE FROM chats WHERE id=$1"

	//Delete all anonymous users in the "users" table.
	DELETE_ANONYMOUS_USERS string = "DELETE FROM users WHERE is_anonymous IS TRUE"

	//Delete a user from "users" by their id.
	DELETE_USER string = "DELETE FROM users WHERE id=$1"
)