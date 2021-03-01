package betypes

// NewChats returns *Chats.
//
// The StartAddingUsers function starts working in goroutine.
func NewChats(usersCount, buffer int) *Chats {
	c := &Chats{
		Chats:      make([]Chat, 0),
		UsersCount: usersCount,
		Queue:      make(chan User, buffer),
	}

	go c.StartAddingUsers()

	return c
}

// StartAddingUsers takes the user from the queue.
// Adds him to the chat taking into account the settings that the user specified (gender, age, city).
//
// Blocks the thread in which this function is running,
// since it has an infinite loop. It is recommended to run in goroutine.
func (c *Chats) StartAddingUsers() {
	for u := range c.Queue {
		if len(c.Chats) == 0 {
			c.AddChat(*NewChat())
			c.Chats[len(c.Chats)-1].AddUser(u)
			continue
		}

		chats := c.FindSuitableChatsForUser(u)
		if chats != nil {
			chats[0].AddUser(u)
			continue
		}

		c.AddChat(*NewChat())
		c.Chats[len(c.Chats)-1].AddUser(u)
	}
}

// AddChat adds new chat.
func (c *Chats) AddChat(chat Chat) {
	c.Chats = append(c.Chats, chat)
}

// FindSuitableChatsForUser returns all chats to which the user can be added.
// Takes into account the settings specified by the user (gender, age, city).
//
// Returns a nil if no chat has been matched.
func (c *Chats) FindSuitableChatsForUser(user User) []*Chat {
	chats := make([]*Chat, 0)
	for i := 0; i < len(c.Chats); i++ {
		usersCount := len(c.Chats[i].Users)
		if usersCount == c.UsersCount {
			continue
		}

		u := c.Chats[i].Users[len(c.Chats[i].Users)-1]
		if user.IsSuitableAge(u) && user.IsSuitableCity(u) && user.IsSuitableSex(u) {
			chats = append(chats, &c.Chats[i])
		}
	}

	if len(chats) != 0 {
		return chats
	}

	return nil
}

// AddUserToQueue adds a user to the queue.
func (c *Chats) AddUserToQueue(user User) {
	c.Queue <- user
}

// DeleteUserFromChat removes the user from the chat he is in.
//
// If the user is not in the chat does not do anything.
func (c *Chats) DeleteUserFromChat(userID int) {
	for i := 0; i < len(c.Chats); i++ {
		for j := 0; j < len(c.Chats[i].Users); j++ {
			if c.Chats[i].Users[j].ID == userID {
				c.Chats[i].Users[len(c.Chats[i].Users)-1],
					c.Chats[i].Users[i] = c.Chats[i].Users[i], c.Chats[i].Users[len(c.Chats[i].Users)-1]
				c.Chats[i].Users = c.Chats[i].Users[:len(c.Chats[i].Users)-1]
				return
			}
		}
	}
}

// DeleteChatWithUser deletes the chat that the user is in.
//
// If the user is not in the chat does not do anything.
func (c *Chats) DeleteChatWithUser(userID int) {
	for i := 0; i < len(c.Chats); i++ {
		for j := 0; j < len(c.Chats[i].Users); j++ {
			if c.Chats[i].Users[j].ID == userID {
				c.Chats[len(c.Chats)-1], c.Chats[i] = c.Chats[i], c.Chats[len(c.Chats)-1]
				c.Chats = c.Chats[:len(c.Chats)-1]
				return
			}
		}
	}
}

// GetInterlocutorsByUserID returns all interlocutors by user ID.
//
// If the user is not in the chat, or there are no interlocutors yet, returns nil.
func (c *Chats) GetInterlocutorsByUserID(userID int) []User {
	if !c.IsUserInChat(userID) {
		return nil
	}

	for i := 0; i < len(c.Chats); i++ {
		if c.Chats[i].IsUserInChat(userID) {
			return c.Chats[i].GetInterlocutorsByUserID(userID)
		}
	}

	return nil
}

// IsUserInChat returns true if the user is in the chat.
func (c *Chats) IsUserInChat(userID int) bool {
	for i := 0; i < len(c.Chats); i++ {
		if c.Chats[i].IsUserInChat(userID) {
			return true
		}
	}
	return false
}
