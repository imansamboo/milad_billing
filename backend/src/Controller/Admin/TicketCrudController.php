<?php

namespace App\Controller\Admin;

use App\Entity\Ticket;
use App\Entity\User;
use EasyCorp\Bundle\EasyAdminBundle\Controller\AbstractCrudController;
use EasyCorp\Bundle\EasyAdminBundle\Field\ChoiceField;
use EasyCorp\Bundle\EasyAdminBundle\Field\DateTimeField;
use EasyCorp\Bundle\EasyAdminBundle\Field\Field;
use EasyCorp\Bundle\EasyAdminBundle\Field\IntegerField;
use EasyCorp\Bundle\EasyAdminBundle\Field\MoneyField;
use EasyCorp\Bundle\EasyAdminBundle\Field\TextEditorField;
use EasyCorp\Bundle\EasyAdminBundle\Field\TextField;
use EasyCorp\Bundle\EasyAdminBundle\Field\IdField;


class TicketCrudController extends AbstractCrudController
{
    public static function getEntityFqcn(): string
    {
        return Ticket::class;
    }


    public function configureFields(string $pageName): iterable
    {
        $subject = TextField::new('subject');
        $status = IntegerField::new('status');
        $creator = ChoiceField::new("creator")
            ->setChoices($this->getDoctrine()->getRepository(User::class)->getUserChoices())
            ->onlyOnForms();
        $first_creator = Field::new("firstCreator", "Creator")->onlyOnIndex();
        return [
            $subject,
            $status,
            $creator,
            $first_creator
        ];
    }

}
