<?php

namespace App\Controller\Admin;

use App\Entity\TicketPost;
use EasyCorp\Bundle\EasyAdminBundle\Controller\AbstractCrudController;

class TicketPostCrudController extends AbstractCrudController
{
    public static function getEntityFqcn(): string
    {
        return TicketPost::class;
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
