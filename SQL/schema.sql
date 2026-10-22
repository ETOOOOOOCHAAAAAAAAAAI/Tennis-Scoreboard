CREATE TABLE IF NOT EXISTS Players (
                                       ID SERIAL PRIMARY KEY,
                                       Name VARCHAR(255) NOT NULL UNIQUE
    );

CREATE TABLE IF NOT EXISTS Matches (
                                       ID SERIAL PRIMARY KEY,
                                       Player1_ID INT NOT NULL,
                                       Player2_ID INT NOT NULL,
                                       Winner_ID INT NOT NULL,

                                       CONSTRAINT fk_player1 FOREIGN KEY(Player1_ID) REFERENCES Players(ID),
    CONSTRAINT fk_player2 FOREIGN KEY(Player2_ID) REFERENCES Players(ID),
    CONSTRAINT fk_winner  FOREIGN KEY(Winner_ID)  REFERENCES Players(ID)
    );

CREATE INDEX idx_matches_winner ON Matches(Winner_ID);