<?php

namespace App\Repository;

use App\Entity\Cdr;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @method Cdr|null find($id, $lockMode = null, $lockVersion = null)
 * @method Cdr|null findOneBy(array $criteria, array $orderBy = null)
 * @method Cdr[]    findAll()
 * @method Cdr[]    findBy(array $criteria, array $orderBy = null, $limit = null, $offset = null)
 */
class CdrRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Cdr::class);
    }

    // /**
    //  * @return Cdr[] Returns an array of Cdr objects
    //  */
    /*
    public function findByExampleField($value)
    {
        return $this->createQueryBuilder('c')
            ->andWhere('c.exampleField = :val')
            ->setParameter('val', $value)
            ->orderBy('c.id', 'ASC')
            ->setMaxResults(10)
            ->getQuery()
            ->getResult()
        ;
    }
    */

    /*
    public function findOneBySomeField($value): ?Cdr
    {
        return $this->createQueryBuilder('c')
            ->andWhere('c.exampleField = :val')
            ->setParameter('val', $value)
            ->getQuery()
            ->getOneOrNullResult()
        ;
    }
    */
}
