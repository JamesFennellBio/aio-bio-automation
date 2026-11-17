login to app 
    using an argon2id with 1gb ram reserved, full data encryption using aes-gcm or chacha. 
non-repudiation:
    everything is blake3 hashed to create a constant log of every action adn the state hash. Edb is every time something (ANYthing) changes, a hash is created and a log item created in a sqllite db. Db is encrypted on disk and only decryptedinto ram, it is written to encrypted.
Recovery:
    2-3 shamir's
Kex and transmissions:
    ML-KEM and X25519 are used together. X25519 creates SS1 and ML-KEM SS2. 
        SS1 || SS2 = SS3
        BLAKE3 (SS3) = eK
        Tt(eK)||At --> the wire
        the wire --> TT(eK)||At
        Tt(eK)AT/eK = Tt||At
        Tt recieved, At authenticates
Nonce derivation:
    Nonce will be derived from a 2 byte counter in the packet, big-endian with the direction bit as the final bit of the number, (15 bits for 16 bits)
    
    
