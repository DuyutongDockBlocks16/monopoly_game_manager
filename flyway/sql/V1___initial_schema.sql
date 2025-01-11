CREATE TABLE properties (
    id SERIAL PRIMARY KEY,
    country_name TEXT NOT NULL,
    land_price INTEGER DEFAULT 0,
    house_price INTEGER DEFAULT 0,
    collateral_value INTEGER DEFAULT 0,
    empty_charge INTEGER DEFAULT 0,
    one_house_charge INTEGER DEFAULT 0,
    two_house_charge INTEGER DEFAULT 0,
    hotel_charge INTEGER DEFAULT 0,
    zone_id INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE properties_in_game (
    id SERIAL PRIMARY KEY,
    property_id INTEGER REFERENCES properties(id),
    game_id TEXT NOT NULL,
    house_number INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP    
);


CREATE INDEX properties_id_idx ON properties (id);
CREATE INDEX properties_in_game_game_id_idx ON properties_in_game (game_id);
