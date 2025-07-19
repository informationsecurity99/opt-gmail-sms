CREATE TYPE user_status_enum AS ENUM ('active', 'deleted', 'blocked');
CREATE TYPE otp_status_enum AS ENUM ('unconfirmed', 'confirmed');

-- 2. Table: users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    status user_status_enum NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by UUID
);

-- 3. Table: otp
CREATE TABLE otp (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    status otp_status_enum NOT NULL,
    code VARCHAR(6) NOT NULL,
    expires_at TIMESTAMP
);

-- 4. Table: sysusers
CREATE TABLE sysusers (
    id UUID PRIMARY KEY,
    status user_status_enum NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by UUID
);

-- 5. Table: roles
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    status user_status_enum NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by UUID
);

-- 6. Table: sysuser_roles (many-to-many)
CREATE TABLE sysuser_roles (
    id UUID PRIMARY KEY,
    sysuser_id UUID NOT NULL REFERENCES sysusers(id) ON DELETE CASCADE,
    role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE
);
