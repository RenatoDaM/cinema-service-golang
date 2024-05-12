CREATE TABLE cinema_session (
    id SERIAL PRIMARY KEY,
    film_name VARCHAR(255) NOT NULL,
    session_number INT NOT NULL,
    session_time TIMESTAMP NOT NULL,
    available_seats INT NOT NULL
);

CREATE TABLE ticket (
    id SERIAL PRIMARY KEY,
    cinema_session_id INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    user_id INT NOT NULL,
    seat_number VARCHAR(10) NOT NULL,
    FOREIGN KEY (cinema_session_id) REFERENCES cinema_session(id)
);