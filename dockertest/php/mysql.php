<?php
$dsn = sprintf('mysql:host=%s:3306;dbname=%s',  'mysql', 'newscheme');
$user = 'root';
$password = 'abcdef';
$dbh = new PDO($dsn, $user, $password);
$sql = "SELECT version();";
foreach ($dbh->query($sql, PDO::FETCH_ASSOC) as $row) {
    print_r($row);
    echo '<br/>';
}
$sql = "show tables;";
foreach ($dbh->query($sql, PDO::FETCH_ASSOC) as $row) {
    print_r($row);
    echo '<br/>';
}
echo getenv('HTTP_USER_AGENT');
echo getenv('MYSQL_ROOT_PASSWORD');
echo $_ENV['HTTP_USER_AGENT'];
?>
