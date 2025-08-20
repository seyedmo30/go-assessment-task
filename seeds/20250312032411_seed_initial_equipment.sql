-- +goose Up
-- +goose StatementBegin
INSERT INTO equipment (id, name, stock) VALUES
       (1, 'Robert Julliat 1200 MSI Korrigan Volgspot', 15),
       (2, 'Shure SM58 Microfoon', 15),
       (3, 'Drive-in show  zonder licht ', 6),
       (4, 'Robert Juliat 2500 MSR Heloise', 7),
       (5, 'Stelstok uitschuifbaar 8m', 17),
       (6, 'ETC S4 750W Profielspot 19°', 2),
       (7, 'ETC S4 750W Profielspot 36°', 2),
       (8, 'Selecon Fresnel 1KW', 2),
       (9, 'Selecon Fresnel 2KW', 5),
       (10, 'Par 64, 1000W fourbar MFL Zwart', 9),
       (11, 'Martin Mac 250 Entour moving head', 4),
       (12, 'Martin Mac 575 Krypton Moving head', 1),
       (13, 'Martin Mac 2000 Performance Moving head', 13),
       (14, 'Clay Paky Sharpy', 2),
       (15, 'Eigenfabrikaat Rain cover', 11),
       (16, 'Vliegbeugel 112P', 11),
       (17, 'L-Acoustics - Kiva riggingbumper', 19),
       (18, 'L-Acoustics - Kiva Line-array qr', 2),
       (19, 'L-Acoustics - 112P Set 2x', 15),
       (20, 'DB-Technologies DVX subtop set Medium', 10),
       (21, 'DB-Technologies DVX subtop set Large', 3),
       (22, 'L-Acoustics Monitorset Single', 6),
       (23, 'L-Acoustics Monitorset Band', 19),
       (24, 'Shure Beta 52A', 8),
       (25, 'Shure SM 57', 2),
       (26, 'Shure Beta 87A', 9),
       (27, 'Sennheiser e604 clip', 1),
       (28, 'K&M Microfoonstatief', 16),
       (29, 'K&M 12x Microfoonstatief in case', 17),
       (30, 'Rode NT 5 Overhead', 19),
       (31, 'Microfoonkit Drum', 2),
       (32, 'Microfoonkit Band (set in set)', 12),
       (33, 'BSS actieve Di Box', 14),
       (34, 'Shure UR4D+ tweeweg ontvanger', 13),
       (35, 'Shure UR4D+ vierweg ontvanger', 5),
       (36, 'Shure UR2 Handheld SM58', 5),
       (37, 'Shure UR2 Handheld Beta 87', 9),
       (38, 'Shure UR1M Bodypack zender', 10),
       (39, 'DPA 4060 Headset', 6),
       (40, 'Shure UR4D set 2x SM58 + ontvanger', 17),
       (41, 'Shure UR4D set 2x SM87 + ontvanger', 7),
       (42, 'Drive-in show  met licht', 6),
       (43, 'Pennenkoffer', 7),
       (44, 'Eurotruss trusspen', 17),
       (45, 'Trusshamer', 6);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM equipment WHERE TRUE;
-- +goose StatementEnd
