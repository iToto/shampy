-- SQL SEEDS

-- user
INSERT INTO public.user (email, created_on, updated_on)
    VALUES ('example@gmail.com', NOW(), NOW());

-- shampoo
INSERT INTO shampoo (brand, created_on, updated_on)
    VALUES ('Head And Shoulders', NOW(), NOW());

-- user_shampoo
INSERT INTO user_shampoo (user_id, shampoo_id, created_on, updated_on)
    VALUES (
        (SELECT id FROM public.user WHERE email = 'example@gmail.com'),
        (SELECT id FROM shampoo WHERE brand = 'Head And Shoulders'),
        NOW(),
        NOW()
    );
