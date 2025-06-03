package handlers

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Username et mot de passe requis", http.StatusBadRequest)
		return
	}

	// Vérifier si l'utilisateur existe déjà
	var existing User
	result := DB.Where("username = ?", creds.Username).First(&existing)
	if result.Error == nil {
		http.Error(w, "Nom d'utilisateur déjà utilisé", http.StatusConflict)
		return
	}

	// Hacher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur de hachage", http.StatusInternalServerError)
		return
	}

	user := User{
		Username: creds.Username,
		Password: string(hashedPassword),
	}

	DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur créé"})
}
