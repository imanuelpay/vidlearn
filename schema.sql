-- Create Database
DROP DATABASE IF EXISTS vidlearn;
CREATE DATABASE IF NOT EXISTS vidlearn;
USE vidlearn;

-- Create Table
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `image_url` text DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `tools` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `tool_url` varchar(200) DEFAULT NULL,
  `image_url` text DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `playlists` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `category_id` int DEFAULT NULL,
  `name` text NOT NULL,
  `thumbnail_url` text DEFAULT NULL,
  `description` text DEFAULT NULL,
  `level` varchar(200) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(200) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(200) DEFAULT NULL,

  FOREIGN KEY(category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS `playlist_tools` (
  `playlist_id` int NOT NULL,
  `tool_id` int NOT NULL,
  PRIMARY KEY(playlist_id, tool_id),
  FOREIGN KEY(playlist_id) REFERENCES playlists(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY(tool_id) REFERENCES tools(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `videos` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `playlist_id` int DEFAULT NULL,
  `title` varchar(200) NOT NULL,
  `video_url` text DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,

  FOREIGN KEY(playlist_id) REFERENCES playlists(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `role` tinyint(1) NOT NULL DEFAULT '1',
  `is_reset` tinyint(1) NOT NULL DEFAULT '99',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `verify_code` varchar(100) DEFAULT NULL,
  `verified_at` timestamp NULL DEFAULT NULL
);

-- Insert Data
INSERT INTO `categories` (`id`, `name`, `image_url`) VALUES
(1, 'Programming', 'https://st.depositphotos.com/1062085/4248/v/600/depositphotos_42480949-stock-illustration-set-of-web-page-coding.jpg'),
(2, 'Networking', 'https://st2.depositphotos.com/3899029/5692/v/600/depositphotos_56924299-stock-illustration-people-communications-flat-icons.jpg'),
(3, 'Accounting', 'https://st2.depositphotos.com/1007566/8440/v/600/depositphotos_84403794-stock-illustration-tax-payment-design.jpg'),
(4, 'Office', 'https://st2.depositphotos.com/3474805/6223/v/600/depositphotos_62239579-stock-illustration-creative-office-work-desk.jpg');

INSERT INTO `tools` (`id`, `name`, `tool_url`, `image_url`) VALUES
(1, 'Adobe Illustrator', 'https://www.adobe.com/sea/products/illustrator/free-trial-download.html', 'https://buildwithangga.com/storage/assets/images/tools/logo_adobe_illustrator.png'),
(2, 'Visual Studio Code', 'https://code.visualstudio.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_vs_code_buildwithangga.png'),
(3, 'InVision App', 'https://invisionapp.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_invisionapp.png'),
(4, 'Maze Design', 'https://maze.design/', 'https://buildwithangga.com/storage/assets/images/tools/logo_maze.png'),
(5, 'Firebase', 'https://firebase.google.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_firebase.png'),
(6, 'Flutter SDK', 'https://flutter.dev/', 'https://buildwithangga.com/storage/assets/images/tools/logo_flutter.png'),
(7, 'Figma', 'https://www.figma.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_figma.png'),
(8, 'Laravel', 'https://laravel.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_laravel.png'),
(9, 'Postman API', 'https://www.postman.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_postman_api.png'),
(10, 'GitHub', 'https://github.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_github_buildwithangga.png'),
(11, 'Tailwind CSS', 'https://tailwindcss.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_tailwind.png'),
(12, 'PHP', 'https://php.net/', 'https://buildwithangga.com/storage/assets/images/tools/logo_php.png'),
(13, 'XAMPP', 'https://www.apachefriends.org/index.html', 'https://buildwithangga.com/storage/assets/images/tools/unnamed.png'),
(14, 'Bootstrap CSS', 'https://getbootstrap.com/', 'https://buildwithangga.com/storage/assets/images/tools/logo_bootstrap.png'),
(15, 'Vue JS', 'https://vuejs.org/', 'https://buildwithangga.com/storage/assets/images/tools/logo_vue.png'),
(16, 'Vue CLI', 'https://cli.vuejs.org/', 'https://buildwithangga.com/storage/assets/images/tools/logo_vue.png');

INSERT INTO `playlists` (`id`, `category_id`, `name`, `thumbnail_url`, `description`, `level`, `created_at`, `created_by`, `updated_at`, `updated_by`) VALUES
(1, 2, 'Dart & Flutter Development Bootcamp: Find House App', 'https://buildwithangga.com/storage/assets/thumbnails/Dart%20&%20Flutter%20Development%20Bootcamp_%20Find%20House%20App%20(1).png', 'Membangun aplikasi iOS dan Android dapat memakan biaya yang sangat mahal sekali untuk sebuah Startup. Oleh karena ini kita dapat menggunakan Flutter SDK untuk membuat aplikasi iOS dan Android hanya dalam sekali koding saja (hemat biaya dan waktu). Di kelas ini kita akan belajar bersama mulai dari dasar untuk membangun aplikasi yang membantu masyarakat dalam mencari kos-kosan di seluruh dunia. Aplikasi ini juga terintegrasi dengan API untuk menyimpan data-data penting seperti info kos, dan lainnya.\r\n\r\nApabila saat ini kamu sedang mempelajari Flutter dari awal, maka kelas ini cocok sekali karena kita juga akan belajar tentang bagaimana mengimplementasikan hasil design di Figma ke bentuk code menggunakan Flutter SDK dan Visual Studio Code serta juga Dart. Silakan bergabung dan kami akan tunggu di kelas.', 'Beginner', '2022-02-10 09:14:18', '', NULL, ''),
(2, 1, 'Full-Stack Flutter Mobile Apps Developer', 'https://class.buildwithangga.com/storage/assets/thumbnails/kelas%20online%20flutter%20apps%20developer%20buildwith%20angga.png', 'Kini aplikasi mobile sudah marak digunakan oleh masyarakat Indonesia. Sebagai sarana yang baik untuk komunikasi, bisnis, social networking, pengembangan diri, bahkan perbankan, aplikasi mobile memungkinkan masyarakat kita memenuhi segala kebutuhan hanya melalui handphone tanpa hambatan jarak dan waktu.\r\n\r\nPelajaran kali ini tentang UI UX Design dan Mobile Apps Development menggunakan Flutter SDK untuk iOS dan Android. Bersama mentor kami, Anda akan dipandu membangun aplikasi pembelian tiket nonton bioskop secara online (TIX/CGV) dari tahap dasar hingga aplikasi siap pakai.\r\n\r\nSelain itu, Figma akan digunakan sebagai design tool, sedangkan Visual Studio Code untuk development. Firebase akan digunakan sebagai tempat penyimpanan data pengguna dan transaksi, lalu API dari TheMovieDB sebagai sumber film. Jadi tunggu apa lagi, ayo kita belajar bersama!', 'Beginner', '2022-02-10 09:36:48', '', NULL, '');

INSERT INTO `playlist_tools` (`playlist_id`, `tool_id`) VALUES
(1, 2),
(1, 5),
(1, 6),
(1, 7),
(2, 2),
(2, 3),
(2, 4),
(2, 5),
(2, 6),
(2, 7);

INSERT INTO `videos` (`id`, `playlist_id`, `title`, `video_url`, `description`, `created_at`, `updated_at`) VALUES
(1, 1, 'Perkenalan Flutter', 'https://www.youtube.com/watch?v=kCfd3_-Rq8U', '', '2022-04-27 10:52:09', NULL),
(2, 1, 'Install Flutter', 'https://www.youtube.com/watch?v=v8sfdjox2Eg', '', '2022-04-27 10:52:09', NULL),
(3, 2, 'Brief Projek', 'https://www.youtube.com/watch?v=W7tLLKEh8kk', '', '2022-04-27 10:52:09', NULL),
(4, 2, 'UX Information Architecture', 'https://www.youtube.com/watch?v=SeoCDQaHYEs', '', '2022-04-27 10:52:09', NULL),
(5, 2, 'UX UserFlow', 'https://www.youtube.com/watch?v=kfJqTqlKHOA', '', '2022-04-27 10:52:09', NULL),
(6, 2, 'UX Wireframe', 'https://www.youtube.com/watch?v=gFMyu5gVBCI', '', '2022-04-27 10:52:09', NULL),
(7, 2, 'Review Wireframe', 'https://www.youtube.com/watch?v=tBQ6UwMXZk4', '', '2022-04-27 10:52:09', NULL);

INSERT INTO `users` (`id`, `name`, `email`, `password`, `role`, `is_reset`, `created_at`, `updated_at`, `verify_code`, `verified_at`) VALUES
(1, 'Imanuel Pay', 'immanuelpay@gmail.com', '$2a$10$jh25jAntryHSnFiTV7UH3.eg/4vbsP8mMd3mg8O5vCo0zagXBfmBy', 99, 99, '2022-05-13 12:33:31', '2022-05-16 05:38:36', 'NYZ4viCeztKhSr7pvfAO2mFM2eWlQkV7QbKec98b6vlSQbxxYSLmBWYyOQVxcHJh', '2022-05-13 12:34:35');
