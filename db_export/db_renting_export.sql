PGDMP         *                 |         
   db_renting    15.2    15.2                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    57908 
   db_renting    DATABASE     �   CREATE DATABASE db_renting WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE db_renting;
                postgres    false            �            1259    57909 	   buildings    TABLE     �  CREATE TABLE public.buildings (
    id integer NOT NULL,
    description text,
    address character varying(255),
    country character varying(70),
    category_id integer,
    guests_num integer,
    rooms_num integer,
    bathrooms_num integer,
    price_day integer,
    avalable_from date,
    avalable_untill date,
    user_id integer,
    imgurl text,
    city character varying(100)
);
    DROP TABLE public.buildings;
       public         heap    postgres    false            �            1259    57912    buildings_id_seq    SEQUENCE     �   CREATE SEQUENCE public.buildings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.buildings_id_seq;
       public          postgres    false    214                       0    0    buildings_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.buildings_id_seq OWNED BY public.buildings.id;
          public          postgres    false    215            �            1259    57922 
   categories    TABLE     \   CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(50)
);
    DROP TABLE public.categories;
       public         heap    postgres    false            �            1259    57921    categories_id_seq    SEQUENCE     �   CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.categories_id_seq;
       public          postgres    false    217                       0    0    categories_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;
          public          postgres    false    216            �            1259    57931    users    TABLE     �   CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50),
    password text,
    email character varying(70),
    phone_num character varying(20)
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    57930    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    219                       0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    218            o           2604    57913    buildings id    DEFAULT     l   ALTER TABLE ONLY public.buildings ALTER COLUMN id SET DEFAULT nextval('public.buildings_id_seq'::regclass);
 ;   ALTER TABLE public.buildings ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    214            p           2604    57925    categories id    DEFAULT     n   ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);
 <   ALTER TABLE public.categories ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            q           2604    57934    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    219    219                      0    57909 	   buildings 
   TABLE DATA           �   COPY public.buildings (id, description, address, country, category_id, guests_num, rooms_num, bathrooms_num, price_day, avalable_from, avalable_untill, user_id, imgurl, city) FROM stdin;
    public          postgres    false    214   �       	          0    57922 
   categories 
   TABLE DATA           .   COPY public.categories (id, name) FROM stdin;
    public          postgres    false    217   p                 0    57931    users 
   TABLE DATA           I   COPY public.users (id, username, password, email, phone_num) FROM stdin;
    public          postgres    false    219                      0    0    buildings_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.buildings_id_seq', 17, true);
          public          postgres    false    215                       0    0    categories_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.categories_id_seq', 17, true);
          public          postgres    false    216                       0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 2, true);
          public          postgres    false    218            s           2606    57929    buildings buildings_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.buildings
    ADD CONSTRAINT buildings_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.buildings DROP CONSTRAINT buildings_pkey;
       public            postgres    false    214            u           2606    57927    categories categories_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.categories DROP CONSTRAINT categories_pkey;
       public            postgres    false    217            w           2606    57938    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    219               �  x����N�V�����7�ΰχJ�
0
t*23�T��m{��8vƇ	��9�b]މO˅o�`�"���_���7��[�e������᦭�kя�-���EQLy�I���H8�-�v�|v�/�kN��Y�Y�&��l���*"BJ��R�
I�����_�5>Y�������<��mO��=O�ju��y|^�?ŧ~����~y�sƲ��p������kZd7/:����w�<��]�y�G���<	��..\�XU��c�t�������mm�
[��8��S8h��XsI"�E"=�,=po͡�>q�?�nߊ��l���Hi�n+�"�I�AqC�s !�b
J�т�Ȩ��:�I���K$z������ra��~s��l�$�_��ҡ�:o�rua��c�;���AFG����n�XiY�3CRBK�,��o�b'8P���d��Uن]�d�j�Cb�=C&��pTqUCG��{ b��Qgr8j�����8j�̻�;,R2l�ƒj��S����䘁�U}�c0�}����<񹪟A(�>�%����!'�=TL�
�a���MT8Fi�5|/_2�����As�Ax�\�nc�66��E�pw�-r�ʍE�5l�U���)#\3�����q�\|Ԕ���Ǯia(\� ��Y)(�ݍ��K[��"=�g?�'P1A�$Z�&��T�:�l�j26yc^g >�ɛ��A�-����2�J$��?ݭ��'
fG��CAHP��BS���]Eq���1G�Ɏ87�bp�Ov�����'q��ωV�3o_����c7�֗L��Ԓ��#g)t��L��^h ^� u��M����������ûe^��2,��~3#�x_2��3A�
�'�UJ��B1�� =�����J��R}WwO�F�o�м��E�}"���y�)��K� a�(h#��pj��I�-���cb�8 k_�S�i�ԅ�U��-�k[����T�So��������G�a���/�BCh�d�D��Z��&>�ڰU��g�*/!%d����,���g�v��Sf0E����#�P՗L\,UJ���a����gG\,�r���Y�#�N\Q��H8�F�|{7"�=��J���I<��k!��n���J���{���A�<#�Z      	   �   x���
�0F���#�/K�n�
�7nB2hh�H&}{���(�X�*��`g罠�K����r������1����x�`O���M΅'��kɞj[�Y�
��ZUE�d�3P5,�2TC�N�|-����?�	*�         l   x�3�LL����T1JT14PI���+7INs��t-u�H�I-��OO)O����7�r)K	OuΨ����)O��LLJ�tH�M����6���0�0546425�����  ��     