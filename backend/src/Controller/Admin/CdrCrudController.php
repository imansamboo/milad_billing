<?php

namespace App\Controller\Admin;

use App\Entity\Cdr;
use EasyCorp\Bundle\EasyAdminBundle\Controller\AbstractCrudController;

class CdrCrudController extends AbstractCrudController
{
    public static function getEntityFqcn(): string
    {
        return Cdr::class;
    }

    /*
    public function configureFields(string $pageName): iterable
    {
        return [
            IdField::new('id'),
            TextField::new('title'),
            TextEditorField::new('description'),
        ];
    }
    */
}
