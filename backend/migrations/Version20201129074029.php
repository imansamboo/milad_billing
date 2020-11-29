<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20201129074029 extends AbstractMigration
{
    public function getDescription() : string
    {
        return '';
    }

    public function up(Schema $schema) : void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE TABLE cdr (id INT AUTO_INCREMENT NOT NULL, sim_number INT DEFAULT NULL, account_number VARCHAR(40) DEFAULT NULL, invoice_number VARCHAR(40) DEFAULT NULL, ocean_region VARCHAR(40) DEFAULT NULL, date DATE DEFAULT NULL, time TIME DEFAULT NULL, originator_number INT DEFAULT NULL, subscriber VARCHAR(40) DEFAULT NULL, destination_number VARCHAR(40) NOT NULL, volume DOUBLE PRECISION NOT NULL, unit VARCHAR(10) NOT NULL, rate VARCHAR(40) DEFAULT NULL, total_charge DOUBLE PRECISION NOT NULL, equipment_type VARCHAR(80) DEFAULT NULL, call_2_call_voice VARCHAR(80) DEFAULT NULL, call_identifier_id INT NOT NULL, originator_country VARCHAR(40) NOT NULL, destination_country VARCHAR(40) NOT NULL, provider VARCHAR(40) NOT NULL, upstream_rate INT NOT NULL, downstream_rate INT NOT NULL, data_session_id VARCHAR(40) NOT NULL, apn VARCHAR(255) NOT NULL, PRIMARY KEY(id)) DEFAULT CHARACTER SET utf8mb4 COLLATE `utf8mb4_unicode_ci` ENGINE = InnoDB');
    }

    public function down(Schema $schema) : void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('DROP TABLE cdr');
    }
}
